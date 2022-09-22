package contract_test

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gopherine/evmos-conploy/contract"
)

// ======================================================================================================
// NOTE: This BDD test is not implemented to cover all edgecases , refer table driven tests for complete
// test coverage, tests here are mainly for showcasing purpose these needs to be enhanced or merged along
// with table driven test ideas to get complete coverage
// ======================================================================================================

var _ = Describe("Testing Contract", func() {
	// It creates a new instance of the gomock controller.
	ctrl := gomock.NewController(GinkgoT())
	defer ctrl.Finish()

	var (
		// Creating a mock instance of the `IBlockchain` interface.
		clientMock = contract.NewMockIBlockchain(ctrl)
		c          = contract.NewContract(clientMock)
	)

	Context("Test Deploy function", func() {
		It("Check successful deployment of contract", func() {
			clientMock.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
			clientMock.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
			clientMock.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(1), nil)
			clientMock.EXPECT().EstimateGas(context.Background(), gomock.Any()).Return(uint64(100), nil)
			clientMock.EXPECT().SendTransaction(context.Background(), gomock.Any()).Return(nil)

			instance, addrHash, txHash, err := c.Deploy()

			Expect(err).To(BeNil())
			Expect(instance).ToNot(BeNil())
			Expect(addrHash).To(Equal("0xdB7d6AB1f17c6b31909aE466702703dAEf9269Cf"))
			Expect(txHash).To(Equal("0x48cc3e57257c690e582516f64320a55d471f52c669eb34a55b9e8ecf6a30128d"))
		})

		It("Check Error while getting Chain ID", func() {
			clientMock.EXPECT().ChainID(context.Background()).Return(nil, errors.New("chain id error"))
			_, _, _, err := c.Deploy()

			Expect(err).ToNot(BeNil())
			Expect(err).Error().To(Equal(errors.New("chain id error")))
		})
	})

	Context("Test Read function", func() {
		var instanceMock *contract.MockIGoldcoin
		//TODO: at this point this is not really needed implementing just for testing purpose
		BeforeEach(func() {
			// Creating a mock instance of the `IGoldcoin` interface.
			instanceMock = contract.NewMockIGoldcoin(ctrl)
		})

		It("Check successful execution of Read function", func() {
			instanceMock.EXPECT().Symbol(&bind.CallOpts{}).Return("GLD", nil)
			instanceMock.EXPECT().BalanceOf(&bind.CallOpts{}, testAddr).Return(testBalance, nil)

			symbol, bal, err := c.Read(instanceMock)

			balCheck := bal.Cmp(testBalance)
			Expect(err).To(BeNil())
			Expect(symbol).To(Equal("GLD"))
			Expect(balCheck).To(Equal(0))
		})

		It("Check Error getting Symbol", func() {
			instanceMock.EXPECT().Symbol(&bind.CallOpts{}).Return("", errors.New("error"))
			_, _, err := c.Read(instanceMock)

			Expect(err).ToNot(BeNil())
		})
	})
})
