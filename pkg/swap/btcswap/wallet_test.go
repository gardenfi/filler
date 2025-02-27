package btcswap_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/catalogfi/blockchain/localnet"
	"github.com/catalogfi/cobi/pkg/swap"
	"github.com/catalogfi/cobi/pkg/swap/btcswap"
	"github.com/fatih/color"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("wallet for bitcoin swap", func() {
	Context("Batch execution", func() {
		It("should succeed when doing eight swaps at the same time", func(ctx context.Context) {
			By("Initialization two wallet")
			aliceWallet, err := NewTestWallet(indexer)
			Expect(err).To(BeNil())
			bobWallet, err := NewTestWallet(indexer)
			Expect(err).To(BeNil())

			By("Funding the wallet")
			txhash1, err := localnet.FundBTC(aliceWallet.Address().EncodeAddress())
			Expect(err).To(BeNil())
			By(fmt.Sprintf("Funding address1 %v , txid = %v", aliceWallet.Address(), txhash1))
			txhash2, err := localnet.FundBTC(bobWallet.Address().EncodeAddress())
			Expect(err).To(BeNil())
			By(fmt.Sprintf("Funding address2 %v , txid = %v", bobWallet.Address(), txhash2))
			time.Sleep(5 * time.Second)

			By("Alice and Bob construct their own swap")
			swaps := 8
			waitBlocks := int64(4)
			aliceSwaps := make([]btcswap.Swap, swaps)
			bobSwaps := make([]btcswap.Swap, swaps)
			secrets := make([][]byte, swaps)
			for i := 0; i < swaps; i++ {
				secrets[i] = localnet.RandomSecret()
				secretHash := sha256.Sum256(secrets[i])
				aliceSwaps[i], err = btcswap.NewSwap(network, aliceWallet.Address(), bobWallet.Address(), 1e6, secretHash[:], waitBlocks)
				Expect(err).To(BeNil())
				bobSwaps[i], err = btcswap.NewSwap(network, bobWallet.Address(), aliceWallet.Address(), 2e6, secretHash[:], waitBlocks)
				Expect(err).To(BeNil())
			}

			By("First round : Alice and Bob initialise 1/4 of the swaps")
			aliceActions1 := []btcswap.ActionItem{}
			for i := 0; i < swaps/4; i++ {
				aliceInit := btcswap.ActionItem{
					Action:     swap.ActionInitiate,
					AtomicSwap: aliceSwaps[i],
				}
				aliceActions1 = append(aliceActions1, aliceInit)
			}
			txhash, err := aliceWallet.BatchExecute(ctx, aliceActions1)
			Expect(err).To(BeNil())
			By(color.GreenString("[1] Alice tx hash = %v", txhash))
			bobActions1 := []btcswap.ActionItem{}
			for i := 0; i < swaps/4; i++ {
				bobInit := btcswap.ActionItem{
					Action:     swap.ActionInitiate,
					AtomicSwap: bobSwaps[i],
				}
				bobActions1 = append(bobActions1, bobInit)
			}
			txhash, err = bobWallet.BatchExecute(ctx, bobActions1)
			Expect(err).To(BeNil())
			By(color.GreenString("[1] Bob tx hash = %v", txhash))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)

			By("Wait for a few blocks to be mined")
			for i := int64(0); i < waitBlocks; i++ {
				Expect(localnet.MineBTCBlock()).Should(Succeed())
			}
			time.Sleep(5 * time.Second)

			By("Second round : Alice and Bob initialise next 1/4 of the swaps")
			aliceActions2 := []btcswap.ActionItem{}
			for i := swaps / 4; i < 2*swaps/4; i++ {
				aliceInit := btcswap.ActionItem{
					Action:     swap.ActionInitiate,
					AtomicSwap: aliceSwaps[i],
				}
				aliceActions2 = append(aliceActions2, aliceInit)
			}
			txhash, err = aliceWallet.BatchExecute(ctx, aliceActions2)
			Expect(err).To(BeNil())
			By(color.GreenString("[2] Alice tx hash = %v", txhash))
			bobActions2 := []btcswap.ActionItem{}
			for i := swaps / 4; i < 2*swaps/4; i++ {
				bobInit := btcswap.ActionItem{
					Action:     swap.ActionInitiate,
					AtomicSwap: bobSwaps[i],
				}
				bobActions2 = append(bobActions2, bobInit)
			}
			txhash, err = bobWallet.BatchExecute(ctx, bobActions2)
			Expect(err).To(BeNil())
			By(color.GreenString("[2] Bob tx hash = %v", txhash))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)

			By("Third round : Alice and Bob redeem the 1/4 swaps just been initiated and initiate the next 1/4 swaps")
			aliceActions3 := []btcswap.ActionItem{}
			for i := swaps / 4; i < 3*swaps/4; i++ {
				if i < 2*swaps/4 {
					aliceActions3 = append(aliceActions3, btcswap.ActionItem{
						Action:     swap.ActionRedeem,
						AtomicSwap: bobSwaps[i],
						Secret:     secrets[i],
					})
				} else {
					aliceActions3 = append(aliceActions3, btcswap.ActionItem{
						Action:     swap.ActionInitiate,
						AtomicSwap: aliceSwaps[i],
					})
				}
			}
			txhash, err = aliceWallet.BatchExecute(ctx, aliceActions3)
			Expect(err).To(BeNil())
			By(color.GreenString("[3] Alice tx hash = %v", txhash))
			bobActions3 := []btcswap.ActionItem{}
			for i := swaps / 4; i < 3*swaps/4; i++ {
				if i < 2*swaps/4 {
					bobActions3 = append(bobActions3, btcswap.ActionItem{
						Action:     swap.ActionRedeem,
						AtomicSwap: aliceSwaps[i],
						Secret:     secrets[i],
					})
				} else {
					bobActions3 = append(bobActions3, btcswap.ActionItem{
						Action:     swap.ActionInitiate,
						AtomicSwap: bobSwaps[i],
					})
				}
			}
			txhash, err = bobWallet.BatchExecute(ctx, bobActions3)
			Expect(err).To(BeNil())
			By(color.GreenString("[3] Bob tx hash = %v", txhash))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)

			By("Fourth round : Doing initiate, redeem and refund at the same time")
			aliceActions4 := []btcswap.ActionItem{}
			for i := 0; i < swaps; i++ {
				if i < swaps/4 {
					aliceActions4 = append(aliceActions4, btcswap.ActionItem{
						Action:     swap.ActionRefund,
						AtomicSwap: aliceSwaps[i],
					})
				} else if i >= 2*swaps/4 && i < 3*swaps/4 {
					aliceActions4 = append(aliceActions4, btcswap.ActionItem{
						Action:     swap.ActionRedeem,
						AtomicSwap: bobSwaps[i],
						Secret:     secrets[i],
					})
				} else if i >= 3*swaps/4 && i < swaps {
					aliceActions4 = append(aliceActions4, btcswap.ActionItem{
						Action:     swap.ActionInitiate,
						AtomicSwap: aliceSwaps[i],
					})
				}
			}
			txhash, err = aliceWallet.BatchExecute(ctx, aliceActions4)
			Expect(err).To(BeNil())
			By(color.GreenString("[4] Alice tx hash = %v", txhash))
			bobActions4 := []btcswap.ActionItem{}
			for i := 0; i < swaps; i++ {
				if i < swaps/4 {
					bobActions4 = append(bobActions4, btcswap.ActionItem{
						Action:     swap.ActionRefund,
						AtomicSwap: bobSwaps[i],
					})
				} else if i >= 2*swaps/4 && i < 3*swaps/4 {
					bobActions4 = append(bobActions4, btcswap.ActionItem{
						Action:     swap.ActionRedeem,
						AtomicSwap: aliceSwaps[i],
						Secret:     secrets[i],
					})
				} else if i >= 3*swaps/4 && i < swaps {
					bobActions4 = append(bobActions4, btcswap.ActionItem{
						Action:     swap.ActionInitiate,
						AtomicSwap: bobSwaps[i],
					})
				}
			}
			txhash, err = bobWallet.BatchExecute(ctx, bobActions4)
			Expect(err).To(BeNil())
			By(color.GreenString("[4] Bob tx hash = %v", txhash))
		})

		It("should succeed", func(ctx context.Context) {
			By("Initialization two wallet")
			aliceWallet, err := NewTestWallet(indexer)
			Expect(err).To(BeNil())
			bobWallet, err := NewTestWallet(indexer)
			Expect(err).To(BeNil())

			By("Funding the wallet")
			txhash1, err := localnet.FundBTC(aliceWallet.Address().EncodeAddress())
			Expect(err).To(BeNil())
			By(fmt.Sprintf("Funding address1 %v , txid = %v", aliceWallet.Address(), txhash1))
			txhash2, err := localnet.FundBTC(bobWallet.Address().EncodeAddress())
			Expect(err).To(BeNil())
			By(fmt.Sprintf("Funding address2 %v , txid = %v", aliceWallet.Address(), txhash2))
			time.Sleep(5 * time.Second)

			By("Alice and Bob construct their own swap")
			waitBlocks := int64(6)
			secret := localnet.RandomSecret()
			secretHash := sha256.Sum256(secret)
			aliceSwap, err := btcswap.NewSwap(network, aliceWallet.Address(), bobWallet.Address(), 1e7, secretHash[:], waitBlocks)
			Expect(err).To(BeNil())
			bobSwap, err := btcswap.NewSwap(network, bobWallet.Address(), aliceWallet.Address(), 2e7, secretHash[:], waitBlocks)
			Expect(err).To(BeNil())

			By("Alice initiates her swap")
			aliceInit := btcswap.ActionItem{
				Action:     swap.ActionInitiate,
				AtomicSwap: aliceSwap,
			}
			initiatedTx, err := aliceWallet.BatchExecute(ctx, []btcswap.ActionItem{aliceInit})
			Expect(err).To(BeNil())
			By(color.GreenString("Alice's swap is initiated in tx %v", initiatedTx))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)

			By("Bob initiates his swap")
			bobInit := btcswap.ActionItem{
				Action:     swap.ActionInitiate,
				AtomicSwap: bobSwap,
			}
			initiatedTx, err = bobWallet.BatchExecute(ctx, []btcswap.ActionItem{bobInit})
			Expect(err).To(BeNil())
			By(color.GreenString("Bob's swap is initiated in tx %v", initiatedTx))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)

			By("Alice redeems Bob's swap and reveal the secret")
			aliceRedeem := btcswap.ActionItem{
				Action:     swap.ActionRedeem,
				AtomicSwap: bobSwap,
				Secret:     secret,
			}
			redeemTx, err := aliceWallet.BatchExecute(ctx, []btcswap.ActionItem{aliceRedeem})
			Expect(err).Should(BeNil())
			By(color.GreenString("Bob's swap is redeemed in tx %v", redeemTx))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)
			redeemed, revealedSecret, err := bobSwap.Redeemed(ctx, indexer)
			Expect(err).To(BeNil())
			Expect(redeemed).Should(BeTrue())
			Expect(bytes.Equal(secret, revealedSecret)).Should(BeTrue())

			By("Bob redeems Alice's swap using the revealed secret")
			bobRedeem := btcswap.ActionItem{
				Action:     swap.ActionRedeem,
				AtomicSwap: aliceSwap,
				Secret:     revealedSecret,
			}
			redeemTx, err = bobWallet.BatchExecute(ctx, []btcswap.ActionItem{bobRedeem})
			Expect(err).Should(BeNil())
			By(color.GreenString("Alice's swap is redeemed in tx %v", redeemTx))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)
		})
	})

	Context("Alice and Bob wants to swap 0.1 BTC for 0.2 BTC", func() {
		It("should succeed", func(ctx context.Context) {
			By("Initialization two wallet")
			aliceWallet, err := NewTestWallet(indexer)
			Expect(err).To(BeNil())
			bobWallet, err := NewTestWallet(indexer)
			Expect(err).To(BeNil())

			By("Funding the wallet")
			txhash1, err := localnet.FundBTC(aliceWallet.Address().EncodeAddress())
			Expect(err).To(BeNil())
			By(fmt.Sprintf("Funding address1 %v , txid = %v", aliceWallet.Address(), txhash1))
			txhash2, err := localnet.FundBTC(bobWallet.Address().EncodeAddress())
			Expect(err).To(BeNil())
			By(fmt.Sprintf("Funding address2 %v , txid = %v", aliceWallet.Address(), txhash2))
			time.Sleep(5 * time.Second)

			By("Alice and Bob construct their own swap")
			waitBlocks := int64(6)
			secret := localnet.RandomSecret()
			secretHash := sha256.Sum256(secret)
			aliceSwap, err := btcswap.NewSwap(network, aliceWallet.Address(), bobWallet.Address(), 1e7, secretHash[:], waitBlocks)
			Expect(err).To(BeNil())
			bobSwap, err := btcswap.NewSwap(network, bobWallet.Address(), aliceWallet.Address(), 2e7, secretHash[:], waitBlocks)
			Expect(err).To(BeNil())

			By("Alice initiates her swap")
			initiatedTx, err := aliceWallet.Initiate(ctx, aliceSwap)
			Expect(err).To(BeNil())
			By(color.GreenString("Alice's swap is initiated in tx %v", initiatedTx))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)

			By("Bob initiates his swap")
			initiatedTx, err = bobWallet.Initiate(ctx, bobSwap)
			Expect(err).To(BeNil())
			By(color.GreenString("Bob's swap is initiated in tx %v", initiatedTx))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)

			By("Alice redeems Bob's swap and reveal the secret")
			redeemTx, err := aliceWallet.Redeem(ctx, bobSwap, secret, aliceWallet.Address().EncodeAddress())
			Expect(err).Should(BeNil())
			By(color.GreenString("Bob's swap is redeemed in tx %v", redeemTx))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)
			redeemed, revealedSecret, err := bobSwap.Redeemed(ctx, indexer)
			Expect(err).To(BeNil())
			Expect(redeemed).Should(BeTrue())
			Expect(bytes.Equal(secret, revealedSecret)).Should(BeTrue())

			By("Bob redeems Alice's swap using the revealed secret")
			redeemTx, err = bobWallet.Redeem(ctx, aliceSwap, revealedSecret, bobWallet.Address().EncodeAddress())
			Expect(err).Should(BeNil())
			By(color.GreenString("Alice's swap is redeemed in tx %v", redeemTx))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)
		})
	})

	Context("Alice wants to refund her money after initiation", func() {
		It("should work without any error", func(ctx context.Context) {
			By("Initialization two wallet")
			aliceWallet, err := NewTestWallet(indexer)
			Expect(err).To(BeNil())
			bobWallet, err := NewTestWallet(indexer)
			Expect(err).To(BeNil())

			By("Funding the wallet")
			txhash1, err := localnet.FundBTC(aliceWallet.Address().EncodeAddress())
			Expect(err).To(BeNil())
			By(fmt.Sprintf("Funding address1 %v , txid = %v", aliceWallet.Address(), txhash1))
			time.Sleep(5 * time.Second)

			By("Alice constructs a new swap")
			amount := int64(1e6)
			secret := localnet.RandomSecret()
			secretHash := sha256.Sum256(secret)
			waitBlocks := int64(3)
			aliceSwap, err := btcswap.NewSwap(network, aliceWallet.Address(), bobWallet.Address(), amount, secretHash[:], waitBlocks)
			Expect(err).To(BeNil())

			By("Alice initiates her swap")
			initiatedTx, err := aliceWallet.Initiate(ctx, aliceSwap)
			Expect(err).To(BeNil())
			By(color.GreenString("Alice's swap is initiated in tx %v", initiatedTx))
			Expect(localnet.MineBTCBlock()).Should(Succeed())
			time.Sleep(5 * time.Second)

			By("Wait for a few blocks to be mined")
			for i := int64(0); i < waitBlocks; i++ {
				Expect(localnet.MineBTCBlock()).Should(Succeed())
			}
			time.Sleep(5 * time.Second)

			By("Alice tries to refund her money")
			refundTx, err := aliceWallet.Refund(ctx, aliceSwap, aliceWallet.Address().EncodeAddress())
			Expect(err).To(BeNil())
			By(color.GreenString("Alice's swap is refunded in tx %v", refundTx))
		})
	})
})
