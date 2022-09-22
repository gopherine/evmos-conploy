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
	os.Setenv("OWNER_PRIVATEKEY", "b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
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
