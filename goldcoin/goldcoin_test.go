package goldcoin_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gopherine/evmos-conploy/goldcoin"
)

//NOTE: This test is only a minor implementation using simulated backend to test
// abigen generated go file using simulated backend, this can be extended to test
// other functions easily

func TestSrc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shared Suite")
}

// Test ERC20 contract deployment
var _ = Describe("Deploy Contracts", func() {
	var gasLimit uint64 = 8000029
	key, _ := crypto.GenerateKey() // nolint: gosec
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	genAlloc := make(core.GenesisAlloc)
	genAlloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(9223372036854775807)}

	sim := backends.NewSimulatedBackend(genAlloc, gasLimit)
	defer sim.Close()

	Context("Deploy Goldcoin", func() {
		It("Check successful deployment of contract", func() {
			_, _, _, err := goldcoin.DeployGoldcoin(auth, sim)

			Expect(err).To(BeNil())
		})
	})
})
