package server

import (
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func initMux(deps *Deps) http.Handler {
	mux := http.NewServeMux()

	names := make([]string, 0, len(deps.Services))
	for _, s := range deps.Services {
		s.RegisterHandler(mux)
		names = append(names, s.Name)
	}

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	// Enable gRPC reflection in development mode
	if deps.Cfg.Env == config.EnvDevelopment {
		mux.Handle(grpcreflect.NewHandlerV1(grpcreflect.NewStaticReflector(names...)))
	}

	c := cors.New(cors.Options{
		AllowedOrigins: deps.Cfg.Server.AllowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"Connect-Protocol-Version",
		},
	})

	return c.Handler(h2c.NewHandler(mux, &http2.Server{}))
}
