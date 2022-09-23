package goldcoin_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gopherine/evmos-conploy/goldcoin"
)

//NOTE: This test is only a minor implementation using simulated backend to test
// abigen generated go file using simulated backend, this can be extended to test
// other functions easily

var (
	testKey, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAddr    = crypto.PubkeyToAddress(testKey.PublicKey)
	testBalance = big.NewInt(2e15)
)

func TestSrc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shared Suite")
}

// Test ERC20 contract deployment and tx
var _ = Describe("Deploy Contracts", func() {
	auth, _ := bind.NewKeyedTransactorWithChainID(testKey, big.NewInt(1337))
	balance := new(big.Int).Mul(big.NewInt(999999999999999), big.NewInt(999999999999999))
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: balance},
	}

	sim := backends.NewSimulatedBackend(gAlloc, 8000000)

	Context("Deploy Goldcoin", func() {
		It("Check successful deployment of contract", func() {

			_, _, _, err := goldcoin.DeployGoldcoin(auth, sim)

			Expect(err).To(BeNil())
		})
	})

	Context("Send Transaction", func() {
		It("Check Successful Transaction", func() {
			sim.Blockchain().InsertChain(types.Blocks{})
			nonce, err := sim.PendingNonceAt(context.Background(), auth.From)
			Expect(err).To(BeNil())

			gasPrice, err := sim.SuggestGasPrice(context.Background())
			Expect(err).To(BeNil())

			toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
			var data []byte
			tx := types.NewTransaction(nonce, toAddress, big.NewInt(1000000000000000000), uint64(21000), gasPrice, data)
			signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testKey)
			Expect(err).To(BeNil())

			err = sim.SendTransaction(context.Background(), signedTx)
			Expect(err).To(BeNil())
			Expect(signedTx.Hash().Hex()).ToNot(Equal(BeNil()))

			// Committing the transaction to the blockchain.
			sim.Commit()

			receipt, err := sim.TransactionReceipt(context.Background(), signedTx.Hash())
			Expect(err).To(BeNil())

			Expect(receipt).ToNot(BeNil())
			Expect(receipt.Status).To(Equal(uint64(1)))
		})

	})
})
