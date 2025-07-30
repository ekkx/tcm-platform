package main

import (
	"log/slog"
	"os"

	"github.com/ekkx/tcmrsv-web/cmd/api/server"
)

func main() {
	if err := server.Run(); err != nil {
		slog.Error("server exited with error", slog.String("err", err.Error()))
		os.Exit(1)
	}
}
