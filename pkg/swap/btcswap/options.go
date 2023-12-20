package btcswap

import "github.com/btcsuite/btcd/chaincfg"

type Options struct {
	Network *chaincfg.Params
	FeeTier string
}

func OptionsMainnet() Options {
	return Options{
		Network: &chaincfg.MainNetParams,
		FeeTier: "high",
	}
}

func OptionsTestnet() Options {
	return Options{
		Network: &chaincfg.TestNet3Params,
		FeeTier: "medium",
	}
}

func OptionsRegression() Options {
	return Options{
		Network: &chaincfg.RegressionNetParams,
		FeeTier: "low",
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
