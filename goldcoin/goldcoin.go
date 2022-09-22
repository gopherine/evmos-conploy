// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goldcoin

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GoldcoinMetaData contains all meta data concerning the Goldcoin contract.
var GoldcoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600881526020017f476f6c64636f696e0000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f474c44000000000000000000000000000000000000000000000000000000000081525081600390816200008f9190620004ee565b508060049081620000a19190620004ee565b505050620000e333620000b9620000e960201b60201c565b600a620000c7919062000765565b620f4240620000d79190620007b6565b620000f260201b60201c565b62000903565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000164576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200015b9062000878565b60405180910390fd5b62000178600083836200026a60201b60201c565b80600260008282546200018c91906200089a565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620001e391906200089a565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200024a9190620008e6565b60405180910390a362000266600083836200026f60201b60201c565b5050565b505050565b505050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002f657607f821691505b6020821081036200030c576200030b620002ae565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003767fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000337565b62000382868362000337565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003cf620003c9620003c3846200039a565b620003a4565b6200039a565b9050919050565b6000819050919050565b620003eb83620003ae565b62000403620003fa82620003d6565b84845462000344565b825550505050565b600090565b6200041a6200040b565b62000427818484620003e0565b505050565b5b818110156200044f576200044360008262000410565b6001810190506200042d565b5050565b601f8211156200049e57620004688162000312565b620004738462000327565b8101602085101562000483578190505b6200049b620004928562000327565b8301826200042c565b50505b505050565b600082821c905092915050565b6000620004c360001984600802620004a3565b1980831691505092915050565b6000620004de8383620004b0565b9150826002028217905092915050565b620004f98262000274565b67ffffffffffffffff8111156200051557620005146200027f565b5b620005218254620002dd565b6200052e82828562000453565b600060209050601f83116001811462000566576000841562000551578287015190505b6200055d8582620004d0565b865550620005cd565b601f198416620005768662000312565b60005b82811015620005a05784890151825560018201915060208501945060208101905062000579565b86831015620005c05784890151620005bc601f891682620004b0565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b600185111562000663578086048111156200063b576200063a620005d5565b5b60018516156200064b5780820291505b80810290506200065b8562000604565b94506200061b565b94509492505050565b6000826200067e576001905062000751565b816200068e576000905062000751565b8160018114620006a75760028114620006b257620006e8565b600191505062000751565b60ff841115620006c757620006c6620005d5565b5b8360020a915084821115620006e157620006e0620005d5565b5b5062000751565b5060208310610133831016604e8410600b8410161715620007225782820a9050838111156200071c576200071b620005d5565b5b62000751565b62000731848484600162000611565b925090508184048111156200074b576200074a620005d5565b5b81810290505b9392505050565b600060ff82169050919050565b600062000772826200039a565b91506200077f8362000758565b9250620007ae7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846200066c565b905092915050565b6000620007c3826200039a565b9150620007d0836200039a565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156200080c576200080b620005d5565b5b828202905092915050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b600062000860601f8362000817565b91506200086d8262000828565b602082019050919050565b60006020820190508181036000830152620008938162000851565b9050919050565b6000620008a7826200039a565b9150620008b4836200039a565b9250828201905080821115620008cf57620008ce620005d5565b5b92915050565b620008e0816200039a565b82525050565b6000602082019050620008fd6000830184620008d5565b92915050565b61123880620009136000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c39190610b15565b60405180910390f35b6100e660048036038101906100e19190610bd0565b610308565b6040516100f39190610c2b565b60405180910390f35b61010461032b565b6040516101119190610c55565b60405180910390f35b610134600480360381019061012f9190610c70565b610335565b6040516101419190610c2b565b60405180910390f35b610152610364565b60405161015f9190610cdf565b60405180910390f35b610182600480360381019061017d9190610bd0565b61036d565b60405161018f9190610c2b565b60405180910390f35b6101b260048036038101906101ad9190610cfa565b6103a4565b6040516101bf9190610c55565b60405180910390f35b6101d06103ec565b6040516101dd9190610b15565b60405180910390f35b61020060048036038101906101fb9190610bd0565b61047e565b60405161020d9190610c2b565b60405180910390f35b610230600480360381019061022b9190610bd0565b6104f5565b60405161023d9190610c2b565b60405180910390f35b610260600480360381019061025b9190610d27565b610518565b60405161026d9190610c55565b60405180910390f35b60606003805461028590610d96565b80601f01602080910402602001604051908101604052809291908181526020018280546102b190610d96565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b5050505050905090565b60008061031361059f565b90506103208185856105a7565b600191505092915050565b6000600254905090565b60008061034061059f565b905061034d858285610770565b6103588585856107fc565b60019150509392505050565b60006012905090565b60008061037861059f565b905061039981858561038a8589610518565b6103949190610df6565b6105a7565b600191505092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546103fb90610d96565b80601f016020809104026020016040519081016040528092919081815260200182805461042790610d96565b80156104745780601f1061044957610100808354040283529160200191610474565b820191906000526020600020905b81548152906001019060200180831161045757829003601f168201915b5050505050905090565b60008061048961059f565b905060006104978286610518565b9050838110156104dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d390610e9c565b60405180910390fd5b6104e982868684036105a7565b60019250505092915050565b60008061050061059f565b905061050d8185856107fc565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610616576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060d90610f2e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610685576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067c90610fc0565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516107639190610c55565b60405180910390a3505050565b600061077c8484610518565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107f657818110156107e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107df9061102c565b60405180910390fd5b6107f584848484036105a7565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361086b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610862906110be565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d190611150565b60405180910390fd5b6108e5838383610a7b565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561096b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610962906111e2565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109fe9190610df6565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a629190610c55565b60405180910390a3610a75848484610a80565b50505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610abf578082015181840152602081019050610aa4565b60008484015250505050565b6000601f19601f8301169050919050565b6000610ae782610a85565b610af18185610a90565b9350610b01818560208601610aa1565b610b0a81610acb565b840191505092915050565b60006020820190508181036000830152610b2f8184610adc565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b6782610b3c565b9050919050565b610b7781610b5c565b8114610b8257600080fd5b50565b600081359050610b9481610b6e565b92915050565b6000819050919050565b610bad81610b9a565b8114610bb857600080fd5b50565b600081359050610bca81610ba4565b92915050565b60008060408385031215610be757610be6610b37565b5b6000610bf585828601610b85565b9250506020610c0685828601610bbb565b9150509250929050565b60008115159050919050565b610c2581610c10565b82525050565b6000602082019050610c406000830184610c1c565b92915050565b610c4f81610b9a565b82525050565b6000602082019050610c6a6000830184610c46565b92915050565b600080600060608486031215610c8957610c88610b37565b5b6000610c9786828701610b85565b9350506020610ca886828701610b85565b9250506040610cb986828701610bbb565b9150509250925092565b600060ff82169050919050565b610cd981610cc3565b82525050565b6000602082019050610cf46000830184610cd0565b92915050565b600060208284031215610d1057610d0f610b37565b5b6000610d1e84828501610b85565b91505092915050565b60008060408385031215610d3e57610d3d610b37565b5b6000610d4c85828601610b85565b9250506020610d5d85828601610b85565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610dae57607f821691505b602082108103610dc157610dc0610d67565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e0182610b9a565b9150610e0c83610b9a565b9250828201905080821115610e2457610e23610dc7565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000610e86602583610a90565b9150610e9182610e2a565b604082019050919050565b60006020820190508181036000830152610eb581610e79565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000610f18602483610a90565b9150610f2382610ebc565b604082019050919050565b60006020820190508181036000830152610f4781610f0b565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000610faa602283610a90565b9150610fb582610f4e565b604082019050919050565b60006020820190508181036000830152610fd981610f9d565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611016601d83610a90565b915061102182610fe0565b602082019050919050565b6000602082019050818103600083015261104581611009565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006110a8602583610a90565b91506110b38261104c565b604082019050919050565b600060208201905081810360008301526110d78161109b565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b600061113a602383610a90565b9150611145826110de565b604082019050919050565b600060208201905081810360008301526111698161112d565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006111cc602683610a90565b91506111d782611170565b604082019050919050565b600060208201905081810360008301526111fb816111bf565b905091905056fea26469706673582212202f7fd1ec4722a55967b0e0b75e80ff4b9adf50f4184d1e3d6d6ad7ada2054d8a64736f6c63430008100033",
}

// GoldcoinABI is the input ABI used to generate the binding from.
// Deprecated: Use GoldcoinMetaData.ABI instead.
var GoldcoinABI = GoldcoinMetaData.ABI

// GoldcoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GoldcoinMetaData.Bin instead.
var GoldcoinBin = GoldcoinMetaData.Bin

// DeployGoldcoin deploys a new Ethereum contract, binding an instance of Goldcoin to it.
func DeployGoldcoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Goldcoin, error) {
	parsed, err := GoldcoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GoldcoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Goldcoin{GoldcoinCaller: GoldcoinCaller{contract: contract}, GoldcoinTransactor: GoldcoinTransactor{contract: contract}, GoldcoinFilterer: GoldcoinFilterer{contract: contract}}, nil
}

// Goldcoin is an auto generated Go binding around an Ethereum contract.
type Goldcoin struct {
	GoldcoinCaller     // Read-only binding to the contract
	GoldcoinTransactor // Write-only binding to the contract
	GoldcoinFilterer   // Log filterer for contract events
}

// GoldcoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoldcoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldcoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoldcoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldcoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoldcoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldcoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoldcoinSession struct {
	Contract     *Goldcoin         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoldcoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoldcoinCallerSession struct {
	Contract *GoldcoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GoldcoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoldcoinTransactorSession struct {
	Contract     *GoldcoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GoldcoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoldcoinRaw struct {
	Contract *Goldcoin // Generic contract binding to access the raw methods on
}

// GoldcoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoldcoinCallerRaw struct {
	Contract *GoldcoinCaller // Generic read-only contract binding to access the raw methods on
}

// GoldcoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoldcoinTransactorRaw struct {
	Contract *GoldcoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoldcoin creates a new instance of Goldcoin, bound to a specific deployed contract.
func NewGoldcoin(address common.Address, backend bind.ContractBackend) (*Goldcoin, error) {
	contract, err := bindGoldcoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Goldcoin{GoldcoinCaller: GoldcoinCaller{contract: contract}, GoldcoinTransactor: GoldcoinTransactor{contract: contract}, GoldcoinFilterer: GoldcoinFilterer{contract: contract}}, nil
}

// NewGoldcoinCaller creates a new read-only instance of Goldcoin, bound to a specific deployed contract.
func NewGoldcoinCaller(address common.Address, caller bind.ContractCaller) (*GoldcoinCaller, error) {
	contract, err := bindGoldcoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoldcoinCaller{contract: contract}, nil
}

// NewGoldcoinTransactor creates a new write-only instance of Goldcoin, bound to a specific deployed contract.
func NewGoldcoinTransactor(address common.Address, transactor bind.ContractTransactor) (*GoldcoinTransactor, error) {
	contract, err := bindGoldcoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoldcoinTransactor{contract: contract}, nil
}

// NewGoldcoinFilterer creates a new log filterer instance of Goldcoin, bound to a specific deployed contract.
func NewGoldcoinFilterer(address common.Address, filterer bind.ContractFilterer) (*GoldcoinFilterer, error) {
	contract, err := bindGoldcoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoldcoinFilterer{contract: contract}, nil
}

// bindGoldcoin binds a generic wrapper to an already deployed contract.
func bindGoldcoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GoldcoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Goldcoin *GoldcoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Goldcoin.Contract.GoldcoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Goldcoin *GoldcoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goldcoin.Contract.GoldcoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Goldcoin *GoldcoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Goldcoin.Contract.GoldcoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Goldcoin *GoldcoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Goldcoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Goldcoin *GoldcoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goldcoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Goldcoin *GoldcoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Goldcoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Goldcoin *GoldcoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Goldcoin.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Goldcoin *GoldcoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Goldcoin.Contract.Allowance(&_Goldcoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Goldcoin *GoldcoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Goldcoin.Contract.Allowance(&_Goldcoin.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Goldcoin *GoldcoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Goldcoin.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Goldcoin *GoldcoinSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Goldcoin.Contract.BalanceOf(&_Goldcoin.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Goldcoin *GoldcoinCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Goldcoin.Contract.BalanceOf(&_Goldcoin.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Goldcoin *GoldcoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Goldcoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Goldcoin *GoldcoinSession) Decimals() (uint8, error) {
	return _Goldcoin.Contract.Decimals(&_Goldcoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Goldcoin *GoldcoinCallerSession) Decimals() (uint8, error) {
	return _Goldcoin.Contract.Decimals(&_Goldcoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Goldcoin *GoldcoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Goldcoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Goldcoin *GoldcoinSession) Name() (string, error) {
	return _Goldcoin.Contract.Name(&_Goldcoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Goldcoin *GoldcoinCallerSession) Name() (string, error) {
	return _Goldcoin.Contract.Name(&_Goldcoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Goldcoin *GoldcoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Goldcoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Goldcoin *GoldcoinSession) Symbol() (string, error) {
	return _Goldcoin.Contract.Symbol(&_Goldcoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Goldcoin *GoldcoinCallerSession) Symbol() (string, error) {
	return _Goldcoin.Contract.Symbol(&_Goldcoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Goldcoin *GoldcoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goldcoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Goldcoin *GoldcoinSession) TotalSupply() (*big.Int, error) {
	return _Goldcoin.Contract.TotalSupply(&_Goldcoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Goldcoin *GoldcoinCallerSession) TotalSupply() (*big.Int, error) {
	return _Goldcoin.Contract.TotalSupply(&_Goldcoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Goldcoin *GoldcoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Goldcoin.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Goldcoin *GoldcoinSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.Approve(&_Goldcoin.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Goldcoin *GoldcoinTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.Approve(&_Goldcoin.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Goldcoin *GoldcoinTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Goldcoin.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Goldcoin *GoldcoinSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.DecreaseAllowance(&_Goldcoin.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Goldcoin *GoldcoinTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.DecreaseAllowance(&_Goldcoin.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Goldcoin *GoldcoinTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Goldcoin.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Goldcoin *GoldcoinSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.IncreaseAllowance(&_Goldcoin.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Goldcoin *GoldcoinTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.IncreaseAllowance(&_Goldcoin.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Goldcoin *GoldcoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Goldcoin.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Goldcoin *GoldcoinSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.Transfer(&_Goldcoin.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Goldcoin *GoldcoinTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.Transfer(&_Goldcoin.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Goldcoin *GoldcoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Goldcoin.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Goldcoin *GoldcoinSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.TransferFrom(&_Goldcoin.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Goldcoin *GoldcoinTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Goldcoin.Contract.TransferFrom(&_Goldcoin.TransactOpts, from, to, amount)
}

// GoldcoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Goldcoin contract.
type GoldcoinApprovalIterator struct {
	Event *GoldcoinApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GoldcoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoldcoinApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GoldcoinApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GoldcoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoldcoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoldcoinApproval represents a Approval event raised by the Goldcoin contract.
type GoldcoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Goldcoin *GoldcoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GoldcoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Goldcoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GoldcoinApprovalIterator{contract: _Goldcoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Goldcoin *GoldcoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GoldcoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Goldcoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoldcoinApproval)
				if err := _Goldcoin.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Goldcoin *GoldcoinFilterer) ParseApproval(log types.Log) (*GoldcoinApproval, error) {
	event := new(GoldcoinApproval)
	if err := _Goldcoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GoldcoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Goldcoin contract.
type GoldcoinTransferIterator struct {
	Event *GoldcoinTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GoldcoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoldcoinTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GoldcoinTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GoldcoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoldcoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoldcoinTransfer represents a Transfer event raised by the Goldcoin contract.
type GoldcoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Goldcoin *GoldcoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GoldcoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Goldcoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GoldcoinTransferIterator{contract: _Goldcoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Goldcoin *GoldcoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GoldcoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Goldcoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoldcoinTransfer)
				if err := _Goldcoin.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Goldcoin *GoldcoinFilterer) ParseTransfer(log types.Log) (*GoldcoinTransfer, error) {
	event := new(GoldcoinTransfer)
	if err := _Goldcoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
