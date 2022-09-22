package main

import (
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
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

	_ = client

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
				log.Print("will contain logic for deploying")
			} else if cCtx.Bool("check") {
				log.Print("Will call and check reciept to confirm deployment")
			} else if cCtx.Bool("read") {
				log.Print("Will read from contract some basic functions to check balance and symbol")
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
