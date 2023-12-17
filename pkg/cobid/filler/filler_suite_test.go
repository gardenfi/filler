package filler_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/catalogfi/blockchain/btc"
	"github.com/catalogfi/blockchain/btc/btctest"
	"github.com/catalogfi/cobi/pkg/swap/btcswap"
	"github.com/catalogfi/cobi/pkg/swap/ethswap/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

func TestFiller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filler Suite")
}

var (
	swapAddr  common.Address
	tokenAddr common.Address
	server    *TestOrderBookServer
	Cancel    context.CancelFunc
)

var _ = BeforeSuite(func() {
	By("Required envs")
	Expect(os.Getenv("ETH_URL")).ShouldNot(BeEmpty())
	Expect(os.Getenv("ETH_KEY_1")).ShouldNot(BeEmpty())

	By("Initialise client")
	url := os.Getenv("ETH_URL")
	client, err := ethclient.Dial(url)
	Expect(err).Should(BeNil())
	chainID, err := client.ChainID(context.Background())
	Expect(err).Should(BeNil())

	By("Initialise transactor")
	keyStr := strings.TrimPrefix(os.Getenv("ETH_KEY_1"), "0x")
	keyBytes, err := hex.DecodeString(keyStr)
	Expect(err).Should(BeNil())
	key, err := crypto.ToECDSA(keyBytes)
	Expect(err).Should(BeNil())
	transactor, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	Expect(err).Should(BeNil())

	By("Deploy ERC20 contract")
	var tx *types.Transaction
	tokenAddr, tx, _, err = bindings.DeployTestERC20(transactor, client)
	Expect(err).Should(BeNil())
	_, err = bind.WaitMined(context.Background(), client, tx)
	Expect(err).Should(BeNil())
	By(color.GreenString("ERC20 deployed to %v", tokenAddr.Hex()))

	By("Deploy atomic swap contract")
	swapAddr, tx, _, err = bindings.DeployAtomicSwap(transactor, client, tokenAddr)
	Expect(err).Should(BeNil())
	_, err = bind.WaitMined(context.Background(), client, tx)
	Expect(err).Should(BeNil())
	By(color.GreenString("Atomic swap deployed to %v", swapAddr.Hex()))

	var ctx context.Context
	logger, err := zap.NewDevelopment()
	Expect(err).To(BeNil())
	ctx, Cancel = context.WithCancel(context.Background())
	server = NewTestServer(logger)
	go func() {
		server.Run(ctx, ":8080")
	}()
})

var _ = AfterSuite(func() {
	Cancel()
	fmt.Println("Server Stopped")
})

func NewTestWallet(network *chaincfg.Params, client btc.IndexerClient) (btcswap.Wallet, error) {
	key, _, err := btctest.NewBtcKey(network)
	if err != nil {
		return nil, err
	}
	opts := btcswap.OptionsRegression()
	fee := rand.Intn(18) + 2
	feeEstimator := btc.NewFixFeeEstimator(fee)
	return btcswap.New(opts, client, key, feeEstimator)
}

type TestOrderBookServer struct {
	router *gin.Engine
	logger *zap.Logger
	Msg    chan interface{}
}

func NewTestServer(logger *zap.Logger) *TestOrderBookServer {
	childLogger := logger.With(zap.String("service", "rest"))
	return &TestOrderBookServer{
		router: gin.Default(),
		logger: childLogger,
		Msg:    make(chan interface{}),
	}
}

func (s *TestOrderBookServer) Run(ctx context.Context, addr string) error {
	s.router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	s.router.GET("/", s.socket())
	service := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	go func() {
		if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		s.logger.Info("stopped")
	}()
	<-ctx.Done()
	return service.Shutdown(ctx)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *TestOrderBookServer) socket() gin.HandlerFunc {
	return func(c *gin.Context) {
		mx := new(sync.RWMutex)
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to upgrade to websocket %v", err)})
			return
		}
		defer func() {
			ws.Close()
		}()

		for resp := range s.Msg {
			mx.Lock()
			err = ws.WriteJSON(map[string]interface{}{
				"type": fmt.Sprintf("%T", resp),
				"msg":  resp,
			})
			mx.Unlock()
			if err != nil {
				s.logger.Debug("failed to write message", zap.Error(err))
				return
			}
		}

	}
}
