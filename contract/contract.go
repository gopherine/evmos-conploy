package contract

import (
	"context"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"

	"github.com/gopherine/evmos-conploy/goldcoin"
)

type Contract struct {
	Client IBlockchain
}

// Below interface is directly refereced from https://github.com/bonedaddy/go-defi/blob/main/utils/blockchain.go
// IBlockchain is a generalized interface for interacting with the ethereum blockchain
// it satisfies all functions required by the ethclient, and simulated backend types.
// This allows you to use ethclient and the simulated backend interchangeably which is
// particularly useful for testing
type IBlockchain interface {
	// ChainID retrieves the current chain ID for transaction replay protection.
	ChainID(ctx context.Context) (*big.Int, error)
	// CodeAt returns the code of the given account. This is needed to differentiate
	// between contract internal errors and the local chain being out of sync.
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)
	// ContractCall executes an Ethereum contract call with the specified data as the
	// input.
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	// PendingCodeAt returns the code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error)
	// PendingCallContract executes an Ethereum contract call against the pending state.
	PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error)
	// PendingNonceAt retrieves the current pending nonce associated with an account.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	// SuggestGasTipCap retrieves the currently suggested 1559 priority fee to allow
	// a timely execution of a transaction.
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	// EstimateGas tries to estimate the gas needed to execute a specific
	// transaction based on the current pending state of the backend blockchain.
	// There is no guarantee that this is the true gas limit requirement as other
	// transactions may be added or removed by miners, but it should provide a basis
	// for setting a reasonable default.
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
	// SendTransaction injects the transaction into the pending pool for execution.
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	// FilterLogs executes a log filter operation, blocking during execution and
	// returning all the results in one batch.
	//
	// TODO(karalabe): Deprecate when the subscription one can return past data too.
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	// SubscribeFilterLogs creates a background log filtering operation, returning
	// a subscription immediately, which can be used to stream the found events.
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	// SubscribeNewHead returns an event subscription for a new header.
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	// TransactionByHash checks the pool of pending transactions in addition to the
	// blockchain. The isPending return value indicates whether the transaction has been
	// mined yet. Note that the transaction may not be part of the canonical chain even if
	// it's not pending.
	TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error)
	// BlockNumber returns the most recent block number
	//BlockNumber(ctx context.Context) (uint64, error)
	// HeaderByNumber returns a block header from the current canonical chain. If
	// number is nil, the latest known header is returned.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

// Defining the interface for the goldcoin contract.
type IGoldcoin interface {
	// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
	Symbol(opts *bind.CallOpts) (string, error)
	// 	BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
	// Solidity: function balanceOf(address account) view returns(uint256)
	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
}

// > The function `NewContract` takes an interface `IBlockchain` as an argument and returns a pointer
// to a `Contract` struct
func NewContract(c IBlockchain) *Contract {
	return &Contract{c}
}

// A variable that is assigned to the function `goldcoin.DeployGoldcoin` : this is for testing purpose, for monkeypatching this variable needs to be public
var Deploy = goldcoin.DeployGoldcoin

// Deploying the contract to the blockchain.
func (c *Contract) Deploy() (instance *goldcoin.Goldcoin, addrHash string, txHash string, err error) {
	auth, err := c.getTxSigner()
	if err != nil {
		return nil, "", "", err
	}

	address, tx, instance, err := Deploy(auth, c.Client)
	if err != nil {
		log.Err(err).Msg("Unable to deploy contract")
		return nil, "", "", err
	}

	// TODO: this return is here only for testing purpose or if it needs to be used globally somehow, eventually needs to be removed or refactored
	return instance, address.Hex(), tx.Hash().Hex(), nil
}

// This function is reading the contract data from the blockchain.
func (c *Contract) Read(instance IGoldcoin) (string, *big.Int, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("OWNER_PRIVATEKEY"))
	if err != nil {
		log.Err(err).Msg("Unable to process privatekey from HEX to ECDSA")
		return "", nil, err
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Err(err).Msg("Unable to get symbol")
		return "", nil, err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Err(err).Msg("unable to get balance")
		return "", nil, err
	}

	return symbol, bal, nil
}

// This function is loading the contract from the blockchain.
func (c *Contract) Load() *goldcoin.Goldcoin {
	address := common.HexToAddress((os.Getenv("CONTRACT_ADDRESS")))
	instance, _ := goldcoin.NewGoldcoin(address, c.Client)
	log.Info().Msg("contract is loaded")
	return instance
}

// This function is reading the contract data from the blockchain.
func (c *Contract) Reciept() (*types.Receipt, error) {
	txHashHex := os.Getenv("CONTRACT_HASH")
	txHash := common.HexToHash(txHashHex)

	reciept, err := c.Client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Err(err).Msg("reciept not recieved, contract was not deployed")
		return nil, err
	}

	return reciept, nil
}

// This function is creating a transaction signer.
func (c *Contract) getTxSigner() (*bind.TransactOpts, error) {
	chainId, err := c.Client.ChainID(context.Background())
	if err != nil {
		log.Err(err).Msg("unable to get chain_id for evmos")
		return nil, err
	}

	// Convert privatekey Hex to ECDSA to create transaction signer
	privateKey, err := crypto.HexToECDSA(os.Getenv("OWNER_PRIVATEKEY"))
	if err != nil {
		log.Err(err).Msg("Unable to generate private key")
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Err(err).Msg("unable to create transaction signer")
		return nil, err
	}

	gasPrice, err := c.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Err(err).Msg("unable to get suggested gas price")
		return nil, err
	}

	nonce, err := c.Client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Err(err).Msg("unable to get nonce")
		return nil, err
	}

	gasLimit, err := c.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: auth.From,
		To:   nil,
		Data: common.FromHex(goldcoin.GoldcoinMetaData.Bin),
	})
	if err != nil {
		log.Err(err).Msg("Unable to estimate gas limit")
		return nil, err
	}

	// Setting the transaction parameters.
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit   // in units
	auth.GasPrice = gasPrice

	return auth, nil
}
