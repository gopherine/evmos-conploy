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
	// goldcoin contract instance
	instance := c.Load()
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
				Name:    "reciept",
				Aliases: []string{"r"},
				Usage:   "Check if smart contract is deployed",
			},
			&cli.BoolFlag{
				Name:    "transact",
				Aliases: []string{"t"},
				Usage:   "Transfer tokens from one address to other",
			},
			&cli.BoolFlag{
				Name:    "balanceOf",
				Aliases: []string{"b"},
				Usage:   "Check balance of given address",
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
			} else if cCtx.Bool("reciept") {
				reciept, err := c.Reciept()
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get reciept")
				}
				log.Info().Msgf("Reciept: %v", reciept)
			} else if cCtx.Bool("transact") {
				if cCtx.NArg() > 1 {
					tx, err := c.TransferTokens(instance, cCtx.Args().Get(1) /*address*/, cCtx.Args().Get(0) /*amount*/)
					if err != nil {
						log.Fatal().Err(err).Msg("Transaction failed")
					}

					log.Info().Msgf("TXHash: %v", tx.Hash().String())
				} else {
					log.Fatal().Err(errors.New("Need more arguments, in the format make run fromAddress toAddress")).Msg("transaction failed")
				}
			} else if cCtx.Bool("balanceOf") {
				if cCtx.NArg() > 0 {
					bal, err := c.CheckBal(instance, cCtx.Args().Get(0))
					if err != nil {
						log.Fatal().Err(err).Msg("Unable to get balance")
					}

					log.Info().Msgf("Balance: %v", bal)
				} else {
					bal, err := c.CheckBal(instance, "")
					if err != nil {
						log.Fatal().Err(err).Msg("Unable to get balance")
					}

					log.Info().Msgf("Balance: %v", bal)
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
