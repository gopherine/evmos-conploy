package contract_test

import (
	"context"
	"errors"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gopherine/evmos-conploy/contract"
)

// ======================================================================================================
// NOTE: This BDD test is not implemented to cover all edgecases , refer table driven tests for complete
// test coverage, tests here are mainly for showcasing purpose these needs to be refactored or merged along
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

	Context("Test CheckBal function", func() {
		var instanceMock *contract.MockIGoldcoin
		//TODO: at this point this is not really needed implementing just for testing purpose
		BeforeEach(func() {
			// Creating a mock instance of the `IGoldcoin` interface.
			instanceMock = contract.NewMockIGoldcoin(ctrl)
		})

		It("Check successful execution of Read function", func() {
			instanceMock.EXPECT().BalanceOf(&bind.CallOpts{}, testAddr).Return(testBalance, nil)

			bal, err := c.CheckBal(instanceMock, "")

			balCheck := bal.Cmp(testBalance)
			Expect(err).To(BeNil())
			Expect(balCheck).To(Equal(0))
		})

		It("Check Error getting Symbol", func() {
			instanceMock.EXPECT().BalanceOf(&bind.CallOpts{}, testAddr).Return(nil, errors.New("error"))
			_, err := c.CheckBal(instanceMock, "")

			Expect(err).ToNot(BeNil())
		})
	})

	Context("Test Reciept function", func() {
		It("Check successful execution of Reciept function", func() {
			const txHex = "0x3a33a98d6eb8d2b0e2a0fd1f4cf9d071992cbb0cc4e0e9887711dde505259e9b"
			os.Setenv("CONTRACT_HASH", txHex)
			defer os.Unsetenv("CONTRACT_HASH")

			clientMock.EXPECT().TransactionReceipt(context.Background(), gomock.Any()).Return(&types.Receipt{
				TxHash: common.HexToHash(txHex),
			}, nil)

			r, err := c.Reciept()

			Expect(err).To(BeNil())
			Expect(r).NotTo(BeNil())
		})

		It("Check Error getting Reciept", func() {
			clientMock.EXPECT().TransactionReceipt(context.Background(), gomock.Any()).Return(nil, errors.New("error"))

			_, err := c.Reciept()

			Expect(err).ToNot(BeNil())
		})
	})

	Context("Test TransferToken function", func() {
		It("Check successful execution of transfertoken function", func() {
			instanceMock := contract.NewMockIGoldcoin(ctrl)

			clientMock.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
			clientMock.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
			clientMock.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(1), nil)
			clientMock.EXPECT().EstimateGas(context.Background(), gomock.Any()).Return(uint64(100), nil)
			instanceMock.EXPECT().Transfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(&types.Transaction{}, nil)

			tx, err := c.TransferTokens(instanceMock, "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d", "100")

			Expect(err).To(BeNil())
			Expect(tx).ToNot(BeNil())
		})

		It("Invalid Amount Error", func() {
			instanceMock := contract.NewMockIGoldcoin(ctrl)

			clientMock.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
			clientMock.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
			clientMock.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(1), nil)
			clientMock.EXPECT().EstimateGas(context.Background(), gomock.Any()).Return(uint64(100), nil)

			_, err := c.TransferTokens(instanceMock, "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d", "XXXX")

			Expect(err).ToNot(BeNil())
		})
	})
})
