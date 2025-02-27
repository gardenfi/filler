package ethswap

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/catalogfi/blockchain/evm/bindings/contracts/htlc/gardenhtlc"
	"github.com/catalogfi/ob/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Swap struct {
	ID         [32]byte
	Initiator  common.Address
	Redeemer   common.Address
	SecretHash common.Hash
	Amount     *big.Int
	Expiry     *big.Int
	Contract   common.Address
}

func NewSwap(initiator, redeemer, contract common.Address, secretHash common.Hash, amount, expiry *big.Int) Swap {
	id := sha256.Sum256(append(secretHash[:], common.BytesToHash(initiator.Bytes()).Bytes()...))

	return Swap{
		ID:         id,
		Initiator:  initiator,
		Redeemer:   redeemer,
		SecretHash: secretHash,
		Amount:     amount,
		Expiry:     expiry,
		Contract:   contract,
	}
}

func (swap *Swap) Initiated(ctx context.Context, client *ethclient.Client) (bool, error) {
	htlc, err := gardenhtlc.NewGardenHTLC(swap.Contract, client)
	if err != nil {
		return false, err
	}
	details, err := htlc.Orders(&bind.CallOpts{Context: ctx}, swap.ID)
	if err != nil {
		return false, err
	}
	return details.InitiatedAt.Uint64() != 0, nil
}

func (swap *Swap) Redeemed(ctx context.Context, client *ethclient.Client) (bool, error) {
	// Check if the swap has been redeemed
	htlc, err := gardenhtlc.NewGardenHTLC(swap.Contract, client)
	if err != nil {
		return false, err
	}
	details, err := htlc.Orders(&bind.CallOpts{Context: ctx}, swap.ID)
	if err != nil {
		return false, err
	}
	return details.IsFulfilled, err
}

func (swap *Swap) Secret(ctx context.Context, client *ethclient.Client, step uint64) ([]byte, error) {
	// Check if the swap has been redeemed
	htlc, err := gardenhtlc.NewGardenHTLC(swap.Contract, client)
	if err != nil {
		return nil, err
	}
	details, err := htlc.Orders(&bind.CallOpts{Context: ctx}, swap.ID)
	if err != nil {
		return nil, err
	}
	if !details.IsFulfilled {
		return nil, fmt.Errorf("swap not redeemed")
	}

	start := details.InitiatedAt
	latestBlock, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	latest := latestBlock.Number()
	if step == 0 {
		step = 500
	}

	// Theoretically people can still redeem after the expiry, but we assume the initiator will refund right after the
	// expiry.
	expiry := big.NewInt(0).Add(details.InitiatedAt, details.Timelock)
	if latest.Cmp(expiry) == 1 {
		latest = expiry
	}

	for start.Cmp(latest) == -1 {
		end := start.Uint64() + step
		if end > latest.Uint64() {
			end = latest.Uint64()
		}
		opts := bind.FilterOpts{
			Start:   start.Uint64(),
			End:     &end,
			Context: ctx,
		}
		iter, err := htlc.FilterRedeemed(&opts, [][32]byte{swap.ID}, [][32]byte{swap.SecretHash})
		if err != nil {
			return nil, err
		}
		if iter.Error() != nil {
			return nil, iter.Error()
		}
		for iter.Next() {
			return iter.Event.Secret, nil
		}
		start = big.NewInt(int64(end))
	}

	return nil, fmt.Errorf("secret not found")
}

func (swap *Swap) Expired(ctx context.Context, client *ethclient.Client) (bool, error) {
	htlc, err := gardenhtlc.NewGardenHTLC(swap.Contract, client)
	if err != nil {
		return false, err
	}

	details, err := htlc.Orders(&bind.CallOpts{Context: ctx}, swap.ID)
	if err != nil {
		return false, err
	}
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		return false, err
	}
	return !details.IsFulfilled && latest-details.InitiatedAt.Uint64() >= details.Timelock.Uint64(), nil
}

func FromAtomicSwap(atomicSwap *model.AtomicSwap) (Swap, error) {
	if atomicSwap.SecretHash == "" {
		return Swap{}, fmt.Errorf("empty secret hash")
	}

	waitBlocks, ok := new(big.Int).SetString(atomicSwap.Timelock, 10)
	if !ok {
		return Swap{}, fmt.Errorf("failed to decode timelock")
	}
	amount, ok := new(big.Int).SetString(atomicSwap.Amount, 10)
	if !ok {
		return Swap{}, fmt.Errorf("failed to decode amount")
	}

	if !common.IsHexAddress(atomicSwap.InitiatorAddress) {
		return Swap{}, fmt.Errorf("failed to decode initiator address")
	}
	initiatorAddr := common.HexToAddress(atomicSwap.InitiatorAddress)

	if !common.IsHexAddress(atomicSwap.RedeemerAddress) {
		return Swap{}, fmt.Errorf("failed to decode redeemer address")
	}
	redeemerAddr := common.HexToAddress(atomicSwap.RedeemerAddress)

	if !common.IsHexAddress(string(atomicSwap.Asset)) {
		return Swap{}, fmt.Errorf("failed to decode asset address")
	}
	contractAddr := common.HexToAddress(string(atomicSwap.Asset))

	return NewSwap(initiatorAddr, redeemerAddr, contractAddr, common.HexToHash(atomicSwap.SecretHash), amount, waitBlocks), nil
}
