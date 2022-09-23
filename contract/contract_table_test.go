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
	"github.com/stretchr/testify/assert"

	"github.com/gopherine/evmos-conploy/contract"
	"github.com/gopherine/evmos-conploy/goldcoin"
)

// `TestDeploy` is a test function that tests the `Deploy` function
func (ts *TableSuite) TestDeploy() {
	// Setting a mock environment variable for testing.
	os.Setenv("OWNER_PRIVATEKEY", testKeyStr)
	defer os.Unsetenv("OWNER_PRIVATEKEY")
	// Mock deploy function that returns a `common.Address`, a `*types.Transaction`, a `*goldcoin.Goldcoin` and an
	// `error`.
	var deployFunc = func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *goldcoin.Goldcoin, error) {
		parsed, err := goldcoin.GoldcoinMetaData.GetAbi()
		if err != nil {
			return common.Address{}, nil, nil, err
		}
		if parsed == nil {
			return common.Address{}, nil, nil, errors.New("GetABI returned nil")
		}

		address, tx, _, err := bind.DeployContract(auth, *parsed, common.FromHex(goldcoin.GoldcoinBin), backend)
		if err != nil {
			return common.Address{}, nil, nil, err
		}
		return address, tx, &goldcoin.Goldcoin{GoldcoinCaller: goldcoin.GoldcoinCaller{}, GoldcoinTransactor: goldcoin.GoldcoinTransactor{}, GoldcoinFilterer: goldcoin.GoldcoinFilterer{}}, nil
	}

	// Creating a slice of structs that contains the name of the test, the deploy function, the prepare
	// function and the wantErr boolean.
	subtests := []struct {
		name    string
		deploy  func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *goldcoin.Goldcoin, error)
		prepare func(m *contract.MockIBlockchain)
		wantErr bool
	}{
		{
			name:   "Deploying smart contract Successfully",
			deploy: deployFunc,
			prepare: func(m *contract.MockIBlockchain) {
				m.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
				m.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
				m.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(1), nil)
				m.EXPECT().EstimateGas(context.Background(), gomock.Any()).Return(uint64(100), nil)
				m.EXPECT().SendTransaction(context.Background(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "Error ChainID",
			deploy: deployFunc,
			prepare: func(m *contract.MockIBlockchain) {
				m.EXPECT().ChainID(context.Background()).Return(nil, errors.New("chain id error"))
			},
			wantErr: true,
		}, {
			name:   "Error Auth",
			deploy: deployFunc,
			prepare: func(m *contract.MockIBlockchain) {
				m.EXPECT().ChainID(context.Background()).Return(nil, nil)
			},
			wantErr: true,
		}, {
			name:   "Error SuggestGasPrice",
			deploy: deployFunc,
			prepare: func(m *contract.MockIBlockchain) {
				m.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
				m.EXPECT().SuggestGasPrice(context.Background()).Return(nil, errors.New("suggest gas price error"))
			},
			wantErr: true,
		},
		{
			name:   "Error PendingNonceAt",
			deploy: deployFunc,
			prepare: func(m *contract.MockIBlockchain) {
				m.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
				m.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
				m.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(0), errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:   "Error EstimateGas",
			deploy: deployFunc,
			prepare: func(m *contract.MockIBlockchain) {
				m.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
				m.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
				m.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(1), nil)
				m.EXPECT().EstimateGas(context.Background(), gomock.Any()).Return(uint64(0), errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:   "Error SendTransaction",
			deploy: deployFunc,
			prepare: func(m *contract.MockIBlockchain) {
				m.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
				m.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
				m.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(1), nil)
				m.EXPECT().EstimateGas(context.Background(), gomock.Any()).Return(uint64(100), nil)
				m.EXPECT().SendTransaction(context.Background(), gomock.Any()).Return(errors.New("error"))
			},
			wantErr: true,
		}, {
			name:   "Error Invalid Private Key",
			deploy: deployFunc,
			prepare: func(m *contract.MockIBlockchain) {
				// overwrite private key env variable for checking HexToECDSA error
				os.Setenv("OWNER_PRIVATEKEY", "invalid_key")
				m.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range subtests {
		ts.Run(tt.name, func() {
			// Setting the `Deploy` function to the `deployFunc` function.
			contract.Deploy = tt.deploy

			if tt.prepare != nil {
				// m.Expect expects only successful calls any subsequent calls post a returning error will panic
				// hence this functions helps in presetting such explicit requirement for each edgecase
				tt.prepare(ts.ClientMock)
			}

			if instance, addrHash, txHash, err := ts.Contract.Deploy(); (err != nil) != tt.wantErr {
				ts.Errorf(err, "Expected %v got %v", tt.wantErr, err)
			} else if !tt.wantErr {
				assert.NotNil(ts.T(), instance)
				assert.Equal(ts.T(), addrHash, "0xdB7d6AB1f17c6b31909aE466702703dAEf9269Cf")
				assert.Equal(ts.T(), txHash, "0x48cc3e57257c690e582516f64320a55d471f52c669eb34a55b9e8ecf6a30128d")
			} else {
				assert.Error(ts.T(), err)
			}
		})
	}
}

// A test function that tests the `CheckBal` function.
func (ts *TableSuite) TestCheckBal() {
	// Setting a mock environment variable for testing.
	os.Setenv("OWNER_PRIVATEKEY", testKeyStr)
	defer os.Unsetenv("OWNER_PRIVATEKEY")

	subtests := []struct {
		name    string
		prepare func(m *contract.MockIGoldcoin)
		wantErr bool
	}{
		{
			name: "Read from instance and get symbol and balance",
			prepare: func(m *contract.MockIGoldcoin) {
				m.EXPECT().BalanceOf(&bind.CallOpts{}, testAddr).Return(testBalance, nil)
			},
			wantErr: false,
		},
		{
			name: "Error Balanceof",
			prepare: func(m *contract.MockIGoldcoin) {
				m.EXPECT().BalanceOf(&bind.CallOpts{}, testAddr).Return(nil, errors.New("error"))
			},
			wantErr: true,
		}, {
			name: "Error Invalid Privatekey",
			prepare: func(m *contract.MockIGoldcoin) {
				os.Setenv("OWNER_PRIVATEKEY", "INVALID_KEY")
			},
			wantErr: true,
		},
	}

	for _, tt := range subtests {
		ts.Run(tt.name, func() {
			if tt.prepare != nil {
				// m.Expect expects only successful calls any subsequent calls post a returning error will panic
				// hence this functions helps in presetting such explicit requirement for each edgecase
				tt.prepare(ts.GoldcoinMock)
			}

			bal, err := ts.Contract.CheckBal(ts.GoldcoinMock, "")

			if (err != nil) != tt.wantErr {
				ts.Errorf(err, "Unexpected Result")
			} else if !tt.wantErr {
				assert.True(ts.T(), bal.Cmp(testBalance) == 0)
			}
		})
	}
}

// A test function that tests the `Load` function.
func (ts *TableSuite) TestLoad() {
	subtests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Loads goldcoin instance",
			wantErr: false,
		},
	}

	for _, tt := range subtests {
		ts.Run(tt.name, func() {
			os.Setenv("CONTRACT_ADDRESS", testAddr.Hex())
			defer os.Unsetenv("CONTRACT_ADDRESS")

			instance := ts.Contract.Load()
			assert.NotNil(ts.T(), instance)
		})
	}
}

// This function is testing the `Reciept` function.
func (ts *TableSuite) TestReciept() {
	const txHex = "0x3a33a98d6eb8d2b0e2a0fd1f4cf9d071992cbb0cc4e0e9887711dde505259e9b"
	os.Setenv("CONTRACT_HASH", txHex)
	defer os.Unsetenv("CONTRACT_HASH")

	subtests := []struct {
		name    string
		prepare func(m *contract.MockIBlockchain)
		wantErr bool
	}{
		{
			name: "Successfully generate reciept",
			prepare: func(m *contract.MockIBlockchain) {
				m.EXPECT().TransactionReceipt(ts.Ctx, gomock.Any()).Return(&types.Receipt{
					TxHash: common.HexToHash(txHex),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "Error generating reciept",
			prepare: func(m *contract.MockIBlockchain) {
				os.Setenv("CONTRACT_HASH", "INVALID_HASH")
				ts.ClientMock.EXPECT().TransactionReceipt(ts.Ctx, gomock.Any()).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range subtests {
		ts.Run(tt.name, func() {
			if tt.prepare != nil {
				tt.prepare(ts.ClientMock)
			}

			r, err := ts.Contract.Reciept()

			if (err != nil) != tt.wantErr {
				ts.Errorf(err, "Unexpected Result")
			} else if !tt.wantErr {
				assert.Equal(ts.T(), common.HexToHash(txHex), r.TxHash)
				ts.NoError(err)
			}
		})
	}
}

func (ts *TableSuite) TestTransferTokens() {
	// Setting a mock environment variable for testing.
	os.Setenv("OWNER_PRIVATEKEY", testKeyStr)
	defer os.Unsetenv("OWNER_PRIVATEKEY")

	subTests := []struct {
		name    string
		wantErr bool
		amount  string
		prepare func(m *contract.MockIBlockchain, mg *contract.MockIGoldcoin)
	}{
		{
			name:    "Transfer token successfully",
			wantErr: false,
			amount:  "100",
			prepare: func(m *contract.MockIBlockchain, mg *contract.MockIGoldcoin) {
				m.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
				m.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
				m.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(1), nil)
				m.EXPECT().EstimateGas(context.Background(), gomock.Any()).Return(uint64(100), nil)
				mg.EXPECT().Transfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(&types.Transaction{}, nil)
			},
		}, {
			name:    "Error transfer",
			wantErr: false,
			amount:  "100",
			prepare: func(m *contract.MockIBlockchain, mg *contract.MockIGoldcoin) {
				m.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
				m.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
				m.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(1), nil)
				m.EXPECT().EstimateGas(context.Background(), gomock.Any()).Return(uint64(100), nil)
				mg.EXPECT().Transfer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
		}, {
			name:    "Error Invalid Amount",
			wantErr: false,
			amount:  "INVALID",
			prepare: func(m *contract.MockIBlockchain, mg *contract.MockIGoldcoin) {
				m.EXPECT().ChainID(context.Background()).Return(big.NewInt(001), nil)
				m.EXPECT().SuggestGasPrice(context.Background()).Return(big.NewInt(1000), nil)
				m.EXPECT().PendingNonceAt(context.Background(), gomock.Any()).Return(uint64(1), nil)
				m.EXPECT().EstimateGas(context.Background(), gomock.Any()).Return(uint64(100), nil)
			},
		}, {
			name:   "Error ChainID",
			amount: "100",
			prepare: func(m *contract.MockIBlockchain, mg *contract.MockIGoldcoin) {
				m.EXPECT().ChainID(context.Background()).Return(nil, errors.New("chain id error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range subTests {
		ts.Run(tt.name, func() {
			if tt.prepare != nil {
				tt.prepare(ts.ClientMock, ts.GoldcoinMock)
			}

			tx, err := ts.Contract.TransferTokens(ts.GoldcoinMock, "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d", tt.amount)

			if (err != nil) != tt.wantErr {
				ts.Errorf(err, "Unexpected Result")
			} else if tt.wantErr {
				assert.Error(ts.T(), err)
			} else {
				assert.NotNil(ts.T(), tx)
				assert.NoError(ts.T(), err)
			}
		})
	}
}
