package main

import (
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/gopherine/evmos-conploy/contract"
)

func main() {
	// Initialize godot  env to load env config
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	// Connect to client
	client, err := ethclient.Dial(os.Getenv("CLIENT_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("Client connection failed")
	} else {
		log.Info().Msg("Client connection successful")
	}

	// initialize deploy contract module
	c := contract.NewContract(client)

	// Creating a CLI app with flags and actions.
	// refer makefile on how to manually trigger these flags
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "deploy",
				Aliases: []string{"d"},
				Usage:   "Deploy smart contract",
			},
			&cli.BoolFlag{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "Check if smart contract is deployed",
			},
			&cli.BoolFlag{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "Read smart contract info: Exec with `make read`",
			},
			&cli.BoolFlag{
				Name:    "transact",
				Aliases: []string{"t"},
				Usage:   `Transfer tokens from one address to other`,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Bool("deploy") {
				_, addrHash, txHash, err := c.Deploy()
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to deploy")
				}
				log.Info().Msgf("Address: %s", addrHash)
				log.Info().Msgf("TXHash: ", txHash)
			} else if cCtx.Bool("check") {
				reciept, err := c.Reciept()
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get reciept")
				}
				log.Info().Msgf("Reciept: %v", reciept)
			} else if cCtx.Bool("read") {
				instance := c.Load()
				symbol, bal, err := c.Read(instance)
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to Read from contract")
				}
				log.Info().Msgf("Balance: %v", bal)
				log.Info().Msgf("Symbol: %s", symbol)
			} else if cCtx.Bool("transact") {
				if cCtx.NArg() > 1 {
					log.Print("Will transfer tokens from one address to other", cCtx.Args().Get(0), cCtx.Args().Get(1))
				} else {
					log.Fatal().Err(errors.New("Need more arguments, in the format make run fromAddress toAddress")).Msg("transaction failed")
				}
			}

			return nil
		},
	}

	// Running the app with the arguments passed in the command line.
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err)
	}
}
