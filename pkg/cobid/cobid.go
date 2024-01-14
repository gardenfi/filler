package cobid

import (
	"encoding/hex"
	"fmt"

	"github.com/catalogfi/blockchain/btc"
	"github.com/catalogfi/cobi/pkg/cobid/executor"
	"github.com/catalogfi/cobi/pkg/cobid/filler"
	"github.com/catalogfi/cobi/pkg/swap/btcswap"
	"github.com/catalogfi/cobi/pkg/swap/ethswap"
	"github.com/catalogfi/cobi/pkg/util"
	"github.com/catalogfi/orderbook/model"
	"github.com/catalogfi/orderbook/rest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type Cobid struct {
	// creator   creator.Creator
	executors executor.Executors
	filler    filler.Filler
}

type BtcChainConfig struct {
	Chain   model.Chain
	Indexer string
}

type EvmChainConfig struct {
	Chain       model.Chain
	SwapAddress string
	URL         string
}

type Config struct {
	Key          string
	OrderbookURL string
	Btc          BtcChainConfig   // chain of the native bitcoin
	Evms         []EvmChainConfig // target evm chains for wbtc
	Strategies   []filler.Strategy
}

func NewCobi(config Config, estimator btc.FeeEstimator) (Cobid, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return Cobid{}, err
	}

	// Decode key
	keyBytes, err := hex.DecodeString(config.Key)
	if err != nil {
		return Cobid{}, err
	}
	key, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return Cobid{}, err
	}

	// Bitcoin wallet and executor
	indexer := btc.NewElectrsIndexerClient(logger, config.Btc.Indexer, btc.DefaultRetryInterval)
	btcWalletOptions := btcswap.NewWalletOptions(config.Btc.Chain.Params())
	btcWallet, err := btcswap.NewWallet(btcWalletOptions, indexer, util.EcdsaToBtcec(key), estimator)
	if err != nil {
		return Cobid{}, err
	}
	storage := executor.NewInMemStore()
	btcExe := executor.NewBitcoinExecutor(config.Btc.Chain, logger, btcWallet, storage)
	executors := []executor.Executor{btcExe}

	// Ethereum wallet and executor
	for _, evm := range config.Evms {
		ethClient, err := ethclient.Dial(evm.URL)
		if err != nil {
			return Cobid{}, err
		}

		swapAddr := common.HexToAddress(evm.SwapAddress)
		ethWalletOptions := ethswap.NewOptions(evm.Chain, swapAddr)
		ethWallet, err := ethswap.NewWallet(ethWalletOptions, key, ethClient)
		if err != nil {
			return Cobid{}, err
		}
		ethExe := executor.NewEvmExecutor(evm.Chain, logger, ethWallet, storage)
		if err != nil {
			return Cobid{}, err
		}
		executors = append(executors, ethExe)
	}
	addr := crypto.PubkeyToAddress(key.PublicKey)
	exes := executor.New(logger, executors, addr.Hex(), storage, config.OrderbookURL)

	// Filler
	client := rest.NewClient(fmt.Sprintf("https://%s", config.OrderbookURL), config.Key)
	token, err := client.Login()
	if err != nil {
		return Cobid{}, err
	}
	if err := client.SetJwt(token); err != nil {
		return Cobid{}, err
	}
	filler := filler.New(config.Strategies, client, config.OrderbookURL, logger)

	return Cobid{
		executors: exes,
		filler:    filler,
	}, nil
}

func (cb Cobid) Start() error {
	cb.executors.Start()
	return cb.filler.Start()
}

func (cb Cobid) Stop() {
	cb.executors.Stop()
	cb.filler.Stop()
}
