package creator_test

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/catalogfi/cobi/pkg/cobid/creator"
	"github.com/catalogfi/orderbook/rest"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MockStore struct{}

func (m MockStore) PutSecret(hash, secret []byte) error {
	return nil
}

var _ = Describe("Creator_setup", Ordered, func() {
	var createStore creator.Store
	BeforeAll(func() {
		createStore = MockStore{}
	})

	Context("Create Orders According to Strategy", func() {
		It("should have created Orders", func() {
			By("Mock a swap contract address")
			swapKey, err := crypto.GenerateKey()
			Expect(err).To(BeNil())
			orderPair := fmt.Sprintf("ethereum_localnet:%s-bitcoin_regtest", crypto.PubkeyToAddress(swapKey.PublicKey))

			sty := creator.NewStrategy(6, 12, new(big.Int).SetInt64(1e6), orderPair, 10)
			orderBookUrl := "localhost:8080"
			cobiKeyStr := strings.TrimPrefix(os.Getenv("ETH_KEY_2"), "0x")
			obRestClient := rest.NewClient("http://"+orderBookUrl, cobiKeyStr)
			jwt, err := obRestClient.Login()
			Expect(err).To(BeNil())
			err = obRestClient.SetJwt(jwt)
			Expect(err).To(BeNil())

			logger, err := zap.NewDevelopment()
			Expect(err).To(BeNil())
			ethKey, err := crypto.GenerateKey()
			Expect(err).To(BeNil())
			btcKey, err := btcec.NewPrivateKey()
			Expect(err).To(BeNil())
			addr, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(btcKey.PubKey().SerializeCompressed()), &chaincfg.RegressionNetParams)
			Expect(err).To(BeNil())
			ctr, err := creator.NewCreator(addr.EncodeAddress(), crypto.PubkeyToAddress(ethKey.PublicKey).Hex(), obRestClient, sty, createStore, logger)
			Expect(err).To(BeNil())
			Expect(ctr.Start()).Should(Succeed())
			defer ctr.Stop()

			// TODO : update the test
			// // sleep for one minute at least 5 orders should have been created at most 10
			// time.Sleep(60 * time.Second)
			// _, err = createStore.OrderByID(5) // read Operation on db
			// Expect(err).To(BeNil())
			//
			// _, err = createStore.OrderByID(10)
			// Expect(err).ToNot(BeNil())
		})
	})
})
