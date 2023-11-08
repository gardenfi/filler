package handlers

import (
	"encoding/hex"
	"fmt"

	"github.com/catalogfi/cobi/wbtc-garden/model"
	"github.com/catalogfi/cobi/wbtc-garden/rest"
	"github.com/catalogfi/cobi/wbtc-garden/swapper/bitcoin"
	"github.com/ethereum/go-ethereum/crypto"
)

func FillOrder(cfg CoreConfig, params RequestFill) error {
	defaultIwStore, _ := bitcoin.NewStore(nil)
	key, err := cfg.Keys.GetKey(model.Ethereum, params.UserAccount, 0)
	if err != nil {
		return fmt.Errorf("Error while getting the signing key: %v", err)
	}
	privKey, err := key.ECDSA()
	if err != nil {
		return fmt.Errorf("Error while getting the private key: %v", err)
	}
	client := rest.NewClient(fmt.Sprintf("https://%s", cfg.EnvConfig.OrderBook), hex.EncodeToString(crypto.FromECDSA(privKey)))
	token, err := client.Login()
	if err != nil {
		return fmt.Errorf("Error while logging in : %v", err)
	}
	if err := client.SetJwt(token); err != nil {
		return fmt.Errorf("Error while setting the JWT: %v", err)
	}
	userStore := cfg.Storage.UserStore(params.UserAccount)

	order, err := client.GetOrder(params.OrderId)
	if err != nil {
		return fmt.Errorf("Error while getting the  order pair: %v", err)
	}

	toChain, fromChain, _, _, err := model.ParseOrderPair(order.OrderPair)
	if err != nil {
		return fmt.Errorf("Error while parsing order pair: %v", err)
	}

	// Get the addresses on different chains.
	fromKey, err := cfg.Keys.GetKey(fromChain, params.UserAccount, 0)
	if err != nil {
		return fmt.Errorf("Error while getting from key: %v", err)
	}
	fromAddress, err := fromKey.Address(fromChain, cfg.EnvConfig.Network, defaultIwStore)
	if err != nil {
		return fmt.Errorf("Error while getting address string: %v", err)
	}
	toKey, err := cfg.Keys.GetKey(toChain, params.UserAccount, 0)
	if err != nil {
		return fmt.Errorf("Error while getting to key: %v", err)
	}
	toAddress, err := toKey.Address(toChain, cfg.EnvConfig.Network, defaultIwStore)
	if err != nil {
		return fmt.Errorf("Error while getting address string: %v", err)
	}

	if err := client.FillOrder(params.OrderId, fromAddress, toAddress); err != nil {
		return fmt.Errorf("Error while filling the order: %v", err)
	}
	if err = userStore.PutSecretHash(order.SecretHash, uint64(params.OrderId)); err != nil {
		return fmt.Errorf("Error while storing the secret hash: %v", err)
	}

	// fmt.Println("Order filled successfully")
	return nil
}
