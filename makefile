# .env.template is a placeholder for env variables
envgen:
	- cat .env.template > .env

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
check:
	- ./bin/conploy -c
read:
	- ./bin/conploy -r
transact:
	- ./bin/conploy -t $(from) $(to)