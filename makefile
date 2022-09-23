# gets private key from local client
OWNER_PRIVATEKEY:=$(shell evmosd keys unsafe-export-eth-key mykey --keyring-backend=test)
# .env.template is a placeholder for env variables
envgen:
	- cat .env.template | sed -e "s/\OWNER_PRIVATEKEY_/${OWNER_PRIVATEKEY}/" > .env

# generate abi, bin and gofile
generate-abi:
	- solc --abi --include-path solidity-contracts/node_modules/ --base-path . solidity-contracts/$(contract).sol -o abi
generate-bin:
	- solc --bin --include-path solidity-contracts/node_modules/ --base-path . solidity-contracts/$(contract).sol -o bin
generate-go:
	- mkdir $(contract)
	- abigen --abi=abi/$(contract).abi --bin=bin/$(contract).bin --pkg=goldcoin --out=goldcoin/$(contract).go

# build application
build:
	- go build -o bin/conploy

# run cli app
deploy:
	- ./bin/conploy -d
reciept:
	- ./bin/conploy -r
balanceOf:
	- ./bin/conploy -b $(address)
transfer:
	- ./bin/conploy -t $(amount) $(to)

# generate mocks
# contract mock
contract-mock:
	- mockgen -package contract -destination contract/contract_mock.go -source contract/contract.go
test-table:
	- go test -v -run ^TestTableSuite$$ ./...
test-bdd:
	- go test -v -run ^TestSrc$$ ./...
test:
	- go test -v ./...