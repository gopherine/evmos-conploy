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

## Testing

## To Do

