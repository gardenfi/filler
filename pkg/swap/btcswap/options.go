package btcswap

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/catalogfi/blockchain/btc"
)

type Options struct {
	Network     *chaincfg.Params
	AddressType btc.AddressType
	FeeTier     string
}

func NewWalletOptions(network *chaincfg.Params) Options {
	switch network.Name {
	case chaincfg.MainNetParams.Name:
		return OptionsMainnet()
	case chaincfg.TestNet3Params.Name:
		return OptionsTestnet()
	case chaincfg.RegressionNetParams.Name:
		return OptionsRegression()
	default:
		panic(fmt.Sprintf("unknown network = %v", network.Name))
	}
}

func OptionsMainnet() Options {
	return Options{
		Network:     &chaincfg.MainNetParams,
		AddressType: btc.AddressP2WPKH,
		FeeTier:     "high",
	}
}

func OptionsTestnet() Options {
	return Options{
		Network:     &chaincfg.TestNet3Params,
		AddressType: btc.AddressP2WPKH,
		FeeTier:     "medium",
	}
}

func OptionsRegression() Options {
	return Options{
		Network:     &chaincfg.RegressionNetParams,
		AddressType: btc.AddressP2WPKH,
		FeeTier:     "low",
	}
}

func (opts Options) WithNetwork(network *chaincfg.Params) Options {
	opts.Network = network
	return opts
}

func (opts Options) WithFeeTier(feeTier string) Options {
	opts.FeeTier = feeTier
	return opts
}

func (opts Options) WithAddressType(addressType btc.AddressType) Options {
	opts.AddressType = addressType
	return opts
}
