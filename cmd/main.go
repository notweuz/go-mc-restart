package main

import (
	"gorestart-minecraft/internal/config"
	"gorestart-minecraft/internal/rcon"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var Version = "25.12.1"

func main() {
	setupLogger()
	log.Info().Msg("Starting go-mc-restart service")

	err := config.LoadConfig("configs/application.yml")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load application configuration")
	}

	rcon.SetupRCONClient(&config.Instance.RconConfig)
	rcon.Client.Connect()
}

func setupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
