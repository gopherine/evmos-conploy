package main

import (
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
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
}
