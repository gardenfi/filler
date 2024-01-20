package executor

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/catalogfi/blockchain/btc"
	"github.com/catalogfi/cobi/pkg/swap"
	"github.com/catalogfi/cobi/pkg/swap/btcswap"
	"github.com/catalogfi/orderbook/model"
	"github.com/catalogfi/orderbook/rest"
	"go.uber.org/zap"
)

type BitcoinExecutor struct {
	chain  model.Chain
	logger *zap.Logger
	wallet btcswap.Wallet
	client rest.Client
	signer string
	store  Store
	stop   chan struct{}
}

func NewBitcoinExecutor(chain model.Chain, logger *zap.Logger, wallet btcswap.Wallet, client rest.Client, store Store, signer string) *BitcoinExecutor {
	exe := &BitcoinExecutor{
		chain:  chain,
		logger: logger,
		wallet: wallet,
		client: client,
		signer: signer,
		store:  store,
		stop:   make(chan struct{}),
	}

	return exe
}

func (be *BitcoinExecutor) Start() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		defer ticker.Stop()

	outer:
		for {
			select {
			case <-ticker.C:
				filter := rest.GetOrdersFilter{
					Taker:   be.signer,
					Verbose: true,
					Status:  int(model.Filled),
				}
				orders, err := be.client.GetOrders(filter)
				if err != nil {
					be.logger.Error("get filled orders", zap.Error(err))
					continue
				}

				rbfOptions, prevOrders, err := be.store.GetRbfInfo()
				if err != nil {
					if !errors.Is(err, ErrNotFound) {
						be.logger.Error("get previous tx info", zap.Error(err))
						continue
					}
				}

				isInitiated := func(swap btcswap.Swap) (bool, error) {
					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()

					initiated, _, err := swap.Initiated(ctx, be.wallet.Indexer())
					return initiated, err
				}

				// Get all the orders we need to execute.
				pendingActions := make([]btcswap.ActionItem, 0, len(orders))
				executedActions := make([]btcswap.ActionItem, 0, len(orders))
				orderIDs := make(map[uint]struct{})
				for _, order := range orders {
					var atomicSwap *model.AtomicSwap
					var action swap.Action
					var executed bool

					iStatus := order.InitiatorAtomicSwap.Status
					fStatus := order.FollowerAtomicSwap.Status

					switch {
					case order.FollowerAtomicSwap.Chain == be.chain:
						// Initiate or Refund
						atomicSwap = order.FollowerAtomicSwap
						if iStatus == model.Initiated && (fStatus == model.Detected || fStatus == model.NotStarted) {
							action = swap.ActionInitiate
							executed = fStatus == model.Detected
						} else if fStatus == model.Expired || fStatus == model.RefundDetected {
							action = swap.ActionRefund
							executed = fStatus == model.RefundDetected
						} else {
							continue
						}
					case order.InitiatorAtomicSwap.Chain == be.chain:
						// Redeem
						if (fStatus == model.Redeemed || fStatus == model.RedeemDetected) &&
							(iStatus == model.Initiated || iStatus == model.RedeemDetected) {
							action = swap.ActionRedeem
							atomicSwap = order.InitiatorAtomicSwap
							executed = iStatus == model.RedeemDetected
							order.InitiatorAtomicSwap.Secret = order.FollowerAtomicSwap.Secret
						} else {
							continue
						}
					default:
						continue
					}

					// Parse the order to an action item
					atomicSwap.SecretHash = order.SecretHash
					btcSwap, err := btcswap.FromAtomicSwap(atomicSwap)
					if err != nil {
						be.logger.Error("failed parse swap", zap.Error(err))
						continue
					}
					var secret []byte
					if action == swap.ActionRedeem {
						secret, err = hex.DecodeString(atomicSwap.Secret)
						if err != nil {
							be.logger.Error("failed decode secret", zap.Error(err))
							continue outer
						}
					}

					// Check if the swap has been initiated before to prevent double initiations.
					if action == swap.ActionInitiate {
						if rbfOptions == nil {
							initiated, err := isInitiated(btcSwap)
							if err != nil {
								be.logger.Error("check swap initiation", zap.Error(err))
								continue
							}
							if initiated {
								continue
							}
						}
					}

					// Add the order to our list
					orderIDs[order.ID] = struct{}{}
					actionItem := btcswap.ActionItem{
						Action:     action,
						AtomicSwap: btcSwap,
						Secret:     secret,
					}
					if executed {
						executedActions = append(executedActions, actionItem)
					} else {
						pendingActions = append(pendingActions, actionItem)
					}
				}

				// Skip if we have no orders to process
				if len(pendingActions) == 0 {
					continue
				}

				// Combine all swap actions and submit a new rbf tx
				pendingActions = append(pendingActions, executedActions...)

				// If any old order is missing from the new order set, that means one of the previous tx has been mined,
				// and we want to start a new series of rbf txs.
				for order := range prevOrders {
					if _, ok := orderIDs[order]; !ok {
						if err := be.store.StoreRbfInfo(nil, nil); err != nil {
							be.logger.Error("store rbf info", zap.Error(err))
						}
						continue outer
					}
				}

				// Skip if no new orders since last time
				if reflect.DeepEqual(prevOrders, orderIDs) {
					continue
				}
				for _, actionItem := range pendingActions {
					log.Printf("%v %v", actionItem.Action, actionItem.AtomicSwap.Address.EncodeAddress())
				}
				log.Printf("rfb = %v", rbfOptions)

				// Submit the transaction
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				tx, newRbf, err := be.wallet.BatchExecute(ctx, pendingActions, rbfOptions)
				cancel()
				if err != nil {
					// Previous tx been included in the block, we wait for orderbook to update the order status and
					// check again later
					if errors.Is(err, btc.ErrTxInputsMissingOrSpent) {
						if err := be.store.StoreRbfInfo(nil, nil); err != nil {
							be.logger.Error("store rbf info", zap.Error(err))
						}
						continue
					}
					be.logger.Error("btc execution", zap.Error(err))
					continue
				}

				// Cache the tx in case we want to replace it in the future
				if err := be.store.StoreRbfInfo(newRbf, orderIDs); err != nil {
					be.logger.Error("store rbf info", zap.Error(err))
				}

				be.logger.Info("✅ [Execution]", zap.String("chain", "btc"), zap.String("txid", tx.TxHash().String()))
			case <-be.stop:
				return
			}
		}
	}()
}

func (be *BitcoinExecutor) Stop() {
	if be.stop != nil {
		close(be.stop)
		be.stop = nil
	}
}
