# .env.template is a placeholder for env variables
envgen:
	- cat .env.template > .env
# run cli app
deploy:
	- go run main.go -d
check:
	- go run main.go -c
read:
	- go run main.go -r

.PHONY: transact
transact:
	- go run main.go -t $(from) $(to)