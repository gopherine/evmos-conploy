package contract_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"

	"github.com/gopherine/evmos-conploy/contract"
)

// `TableSuite` is a struct that contains a `suite.Suite` and a `*backends.SimulatedBackend`.
// @property  - `suite.Suite` is a struct that contains the methods that will be used to run the tests.
type TableSuite struct {
	*suite.Suite
	*contract.Contract
	ClientMock *contract.MockIBlockchain
	Ctrl       *gomock.Controller
	Ctx        context.Context
}

// Creating a test account with a balance of 2e15
var (
	testKey, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAddr    = crypto.PubkeyToAddress(testKey.PublicKey)
	testBalance = big.NewInt(2e15)
)

// Creating a genesis block with a balance of 2e15
var genesis = &core.Genesis{
	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced and accepted by the Ethereum core developers into the Ethash consensus.
	// This configuration is intentionally not using keyed fields to force anyone adding flags to the config to also have to set these fields.
	Config: params.AllEthashProtocolChanges,
	// GenesisAlloc specifies the initial state that is part of the genesis block.
	Alloc:     core.GenesisAlloc{testAddr: {Balance: testBalance}},
	ExtraData: []byte("test genesis"),
	Timestamp: 9000,
	//Initial base fee for EIP-1559 blocks.
	BaseFee: big.NewInt(params.InitialBaseFee),
}

// Suite init for table driven test
func TestTableSuite(t *testing.T) {
	// It creates a new instance of the gomock controller.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Creating a mock instance of the `IBlockchain` interface.
	clientMock := contract.NewMockIBlockchain(ctrl)
	contract := contract.NewContract(clientMock)

	suite.Run(t, &TableSuite{
		Suite:      new(suite.Suite),
		Contract:   contract,
		ClientMock: clientMock,
		Ctx:        context.Background(),
		Ctrl:       ctrl,
	})
}

// Suite init for BDD
func TestSrc(t *testing.T) {
	// Setting a mock environment variable for testing.
	os.Setenv("OWNER_PRIVATEKEY", "b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	defer os.Unsetenv("OWNER_PRIVATEKEY")

	RegisterFailHandler(Fail)
	RunSpecs(t, "Shared Suite")
}
