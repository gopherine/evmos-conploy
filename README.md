# evmos-conploy
A CLI app to deploy smart-contracts on evmos node, and keep track of the contracts deployed.

## Table of Contents
* [Description](#description)
* [Setup](#setup)
* [Usage](#usage)
* [Testing](#testing)
* [To Do](#todo)

## Description

The goal of the project is to create ,deploy and query smart contract. The user should be able to transfer balance, get balance and deployment reciept.
we are currently using go-ethereum, but this should adapt to other libraries in future.

## Setup

- **RESOURCES**
    * https://goethereumbook.org/en/
    * https://ethereum.org/en/developers/docs/programming-languages/golang/
    * https://pkg.go.dev/github.com/ethereum/go-ethereum

- **INSTALLATION**
    * https://docs.evmos.org/validators/quickstart/installation.html
        - I have built the node locally, but you may try to build it via docker, once built run ./init.sh to initialize the node
    * https://geth.ethereum.org/docs/install-and-build/installing-geth
        - install ethereum, this will provide us with tools like abigen and solc to compile solidity and generate go files, abi and bin
    * https://golangci-lint.run/usage/install/ (linter installation)
    * `go install github.com/golang/mock/mockgen@v1.6.0`
        - please check `$GOPATH/go/bin` is updated on your .zshrc or .bashrc file in the path, this will enable us to generate mocks
## Usage
.env.template is a placeholder for env variables, this command will generate .env
```
make envgen
```

To generate abi, bin and gofile run the below make commands respectively
```
make generate-abi
make generate-bin contract=goldcoin
make generate-go contract=goldcoin
```

Build program with `make build` and then trigger cli functions to deploy, check , read or transact and interact
with smartcontracts. Use the below command. The make command is just for
convinience here. If you would like to build and trigger manually you can
do that as well.

```
# Deploy contract
make deploy
# Check if the contract is deployed successfully
make check
# Query smart contract
make read
# Transact tokens from owner_address to reciever_address
make transact from=FROMADDRESS to=RECIEVER_ADDRESS
```

## Testing

We are unit testing using two different patterns i.e Table Driven Tests and Behaviour Driven Tests.
Each of the respective tests are being implemented within their own suites to make it easier to
differentiate and make tests more readable.

To run table driven tests
```
make test-table
```

To run BDD tests
```
make test-bdd
```

Alternatively, if you would like to run both the tests simultaneously you can run that simply
by `go test -v ./...` or using make
```
make test
```

## To Do

