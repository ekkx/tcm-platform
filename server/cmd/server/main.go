package main

import (
	"log"

	"github.com/ekkx/tcmrsv-web/server/cmd/server/commands"
	"github.com/ekkx/tcmrsv-web/server/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed initializing config: %v", err)
	}

	commands.Run(cfg)
}
