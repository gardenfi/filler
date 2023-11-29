package utils

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/catalogfi/cobi/pkg/blockchain"
	"github.com/catalogfi/cobi/pkg/swapper/bitcoin"
	"github.com/catalogfi/cobi/pkg/swapper/ethereum"
	"github.com/catalogfi/cobi/store"
	"github.com/catalogfi/orderbook/model"
	"github.com/catalogfi/orderbook/rest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
)

type Key struct {
	inner *bip32.Key
}

func (key *Key) Interface(chain model.Chain) (interface{}, error) {
	if chain.IsBTC() {
		return key.BtcKey(), nil
	} else {
		return key.ECDSA()
	}
}

func (key *Key) BtcKey() *btcec.PrivateKey {
	privKey, _ := btcec.PrivKeyFromBytes(key.inner.PublicKey().Key)
	return privKey
}

func (key *Key) ECDSA() (*ecdsa.PrivateKey, error) {
	return crypto.ToECDSA(key.inner.Key)
}

func (key *Key) Address(chain model.Chain, network model.Network, isLegacy bool, iwConfig ...bitcoin.InstantWalletConfig) (string, error) {
	switch {
	case chain.IsBTC():
		params := getParams(chain)
		if len(iwConfig) > 0 {
			return key.InstantWalletAddress(chain, network, iwConfig[0])
		}
		var addr btcutil.Address
		var err error
		if isLegacy {
			addr, err = key.LegacyAddress(params)
			if err != nil {
				return "", err
			}
		} else {
			addr, err = key.WitnessAddress(params)
			if err != nil {
				return "", err
			}
		}
		return addr.EncodeAddress(), nil
	case chain.IsEVM():
		addr, err := key.EvmAddress()
		if err != nil {
			return "", err
		}
		return addr.Hex(), nil
	default:
		return "", fmt.Errorf("unsupport chain type %v", chain)
	}
}

func (key *Key) InstantWalletAddress(chain model.Chain, config model.Network, iwConfig bitcoin.InstantWalletConfig) (string, error) {
	client, err := blockchain.LoadClient(chain, config, iwConfig)
	if err != nil {
		return "", fmt.Errorf("failed to load client: %v", err)
	}
	iwClient := client.(bitcoin.InstantClient)
	return iwClient.GetInstantWalletAddress(), nil
}

func (key *Key) LegacyAddress(network *chaincfg.Params) (btcutil.Address, error) {
	keyBytesHash := btcutil.Hash160(key.BtcKey().PubKey().SerializeCompressed())
	return btcutil.NewAddressPubKeyHash(keyBytesHash, network)
}
func (key *Key) WitnessAddress(network *chaincfg.Params) (btcutil.Address, error) {
	keyBytesHash := btcutil.Hash160(key.BtcKey().PubKey().SerializeCompressed())
	return btcutil.NewAddressWitnessPubKeyHash(keyBytesHash, network)
}

func (key *Key) EvmAddress() (common.Address, error) {
	ecdsaKey, err := key.ECDSA()
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(ecdsaKey.PublicKey), nil
}

func LoadKey(seed []byte, chain model.Chain, user, selector uint32) (*Key, error) {
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	var index uint32
	switch chain {
	case model.Bitcoin:
		index = 0
	case model.BitcoinTestnet, model.BitcoinRegtest:
		index = 1
	case model.Ethereum, model.EthereumLocalnet, model.EthereumSepolia, model.EthereumOptimism, model.EthereumArbitrum, model.EthereumAvalanche, model.EthereumBNB, model.EthereumPolygon:
		index = 60
	default:
		return nil, fmt.Errorf("invalid chain: %s", chain)
	}

	for _, idx := range append([]uint32{index}, user, selector) {
		masterKey, err = masterKey.NewChildKey(idx)
		if err != nil {
			return nil, fmt.Errorf("failed to create child key: %v", err)
		}
	}
	return &Key{masterKey}, nil
}

type Keys struct {
	entropy []byte
	m       map[[32]byte]*Key
}

func NewKeys(entropy []byte) Keys {
	return Keys{
		entropy: entropy,
		m:       map[[32]byte]*Key{},
	}
}

func (keys Keys) GetKey(chain model.Chain, user, selector uint32) (*Key, error) {
	digest := append(keys.entropy, []byte(fmt.Sprintf("%v_%v_%v", chain, user, selector))...)
	mapKey := sha256.Sum256(digest)
	value, ok := keys.m[mapKey]
	if !ok {
		var err error
		value, err = LoadKey(keys.entropy, chain, user, selector)
		if err != nil {
			return nil, err
		}
		keys.m[mapKey] = value
	}
	return value, nil
}

func getParams(chain model.Chain) *chaincfg.Params {
	switch chain {
	case model.Bitcoin:
		return &chaincfg.MainNetParams
	case model.BitcoinTestnet:
		return &chaincfg.TestNet3Params
	case model.BitcoinRegtest:
		return &chaincfg.RegressionNetParams
	default:
		panic("constraint violation: unknown chain")
	}
}

func Balance(chain model.Chain, address string, config model.Network, asset model.Asset, iwConfig ...bitcoin.InstantWalletConfig) (*big.Int, error) {
	client, err := blockchain.LoadClient(chain, config, iwConfig...)
	if err != nil {
		return nil, fmt.Errorf("failed to load client: %v", err)
	}

	switch client := client.(type) {

	case bitcoin.Client, bitcoin.InstantClient:
		c := client.(bitcoin.Client)
		address, err := btcutil.DecodeAddress(address, c.Net())
		if err != nil {
			return nil, fmt.Errorf("failed to create address (%s): %v", address, err)
		}
		_, balance, err := c.GetUTXOs(address, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to get UTXOs: %v", err)
		}
		return big.NewInt(int64(balance)), nil

	case ethereum.Client:
		address := common.HexToAddress(address)
		if asset == model.Primary {
			balance, err := client.GetProvider().BalanceAt(context.Background(), address, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to get ETH balance: %v", err)
			}
			return balance, nil
		} else {
			token, err := client.GetTokenAddress(common.HexToAddress(asset.SecondaryID()))
			if err != nil {
				return nil, fmt.Errorf("failed to get ERC20 token address: %v", err)
			}
			balance, err := client.GetERC20Balance(token, address)
			if err != nil {
				return nil, fmt.Errorf("failed to get ERC20 balance: %v", err)
			}
			return balance, nil
		}
	default:
		return nil, fmt.Errorf("unsupported chain: %s", chain)
	}
}

func VirtualBalance(chain model.Chain, address string, config model.Network, asset model.Asset, signer string, client rest.Client, iwConfig ...bitcoin.InstantWalletConfig) (*big.Int, error) {
	var balance *big.Int
	var err error
	balance, err = Balance(chain, address, config, asset, iwConfig...)
	if err != nil {
		return nil, err
	}

	committedAmount := big.NewInt(0)

	fillOrders, err := client.GetOrders(rest.GetOrdersFilter{
		Taker:   signer,
		Status:  int(model.Filled),
		Verbose: true,
	})
	if err != nil {
		return nil, err
	}
	for _, fillOrder := range fillOrders {
		switch fillOrder.Status {
		case model.Created, model.Filled:
		default:
			continue
		}
		if fillOrder.FollowerAtomicSwap.Asset == asset && fillOrder.FollowerAtomicSwap.Status < model.Redeemed {
			orderAmt, ok := new(big.Int).SetString(fillOrder.FollowerAtomicSwap.Amount, 10)
			if !ok {
				return nil, err
			}
			committedAmount.Add(committedAmount, orderAmt)
		}
	}

	// Subtract the amount we open as a maker
	createOrders, err := client.GetOrders(rest.GetOrdersFilter{
		Maker: signer,
		// Status:  int(model.Filled),
		Verbose: true,
	})
	if err != nil {
		return nil, err
	}
	for _, createOrder := range createOrders {
		switch createOrder.Status {
		case model.Created, model.Filled:
		default:
			continue
		}

		if createOrder.InitiatorAtomicSwap.Asset == asset {
			orderAmt, ok := new(big.Int).SetString(createOrder.InitiatorAtomicSwap.Amount, 10)
			if !ok {
				return nil, err
			}
			committedAmount.Add(committedAmount, orderAmt)
		}
	}
	return new(big.Int).Sub(balance, committedAmount), nil
}

func LoadClient(url string, keys Keys, str store.Store, account, selector uint32) (common.Address, rest.Client, error) {
	key, err := keys.GetKey(model.Ethereum, account, selector)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to get the signing key: %v", err)
	}
	privKey, err := key.ECDSA()
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to load ecdsa key: %v", err)
	}
	client := rest.NewClient(fmt.Sprintf("https://%s", url), hex.EncodeToString(crypto.FromECDSA(privKey)))
	signer := crypto.PubkeyToAddress(privKey.PublicKey)

	jwt, err := str.UserStore(account).Token(selector)
	if err != nil {
		jwt, err = client.Login()
		if err != nil {
			return common.Address{}, nil, fmt.Errorf("failed to login to the orderbook: %v", err)
		}
		str.UserStore(account).PutToken(selector, jwt)
	}
	if err := client.SetJwt(jwt); err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to set the jwt token: %v", err)
	}
	return signer, client, nil
}

func ToHex(key *ecdsa.PrivateKey) string {
	return fmt.Sprintf("'%*s'", 10)
}
