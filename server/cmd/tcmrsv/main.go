package main

import (
	"log"

	"github.com/ekkx/tcmrsv-web/server/cmd/tcmrsv/commands"
	"github.com/ekkx/tcmrsv-web/server/pkg/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed initializing config: %v", err)
	}

	commands.Run(cfg)
}
