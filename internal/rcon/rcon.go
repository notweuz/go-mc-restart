package rcon

import (
	"gorestart-minecraft/internal/config"

	"github.com/gorcon/rcon"
	"github.com/rs/zerolog/log"
)

type RCONClient struct {
	rcon   *rcon.Conn
	config *config.RconConfig
}

func SetupRCONClient(config *config.RconConfig) {
	Client = &RCONClient{
		config: config,
	}
}

var Client *RCONClient

func (rc *RCONClient) Connect() error {
	log.Printf("Connecting to %s\n", rc.config.Address)
	conn, err := rcon.Dial(rc.config.Address, rc.config.Password)
	rc.rcon = conn
	return err
}

func (rc *RCONClient) Execute(command string) (string, error) {
	response, err := rc.rcon.Execute(command)
	log.Info().Str("command", command).Str("response", response).Msg("Executed RCON command")
	return response, err
}

func (rc *RCONClient) Close() error {
	err := rc.rcon.Close()
	log.Info().Msg("Closing RCON connection")
	return err
}
