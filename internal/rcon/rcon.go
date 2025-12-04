package rcon

import (
	"gorestart-minecraft/internal/config"

	"github.com/jltobler/go-rcon"
	"github.com/rs/zerolog/log"
)

type RCONClient struct {
	Rcon   *rcon.Client
	Config *config.RconConfig
}

func SetupRCONClient(config *config.RconConfig) {
	Client = &RCONClient{
		Config: config,
	}
}

var Client *RCONClient

func (rc *RCONClient) Connect() {
	log.Printf("Connecting to %s\n", rc.Config.Address)
	rc.Rcon = rcon.NewClient(rc.Config.Address, rc.Config.Password)
}
