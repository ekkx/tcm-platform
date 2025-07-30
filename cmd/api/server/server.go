package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/ekkx/tcmrsv-web/internal/config"
	"github.com/ekkx/tcmrsv-web/internal/shared/logger"
	"github.com/ekkx/tcmrsv-web/pkg/jwt"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	logger.Init(cfg)

	ctx := context.Background()

	dbPool, err := cfg.Database.Open(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbPool.Close()

	jwtManager := jwt.NewJWTManager(
		cfg.Auth.JWTSecretKey,
		cfg.Auth.AccessExpiration,
		cfg.Auth.RefreshExpiration,
	)

	services := getServiceDefinitions(cfg, dbPool, jwtManager)

	mux := http.NewServeMux()

	for _, service := range services {
		service.RegisterHandler(mux)
	}

	// Enable gRPC reflection in development mode
	if cfg.Env == config.EnvDevelopment {
		serviceNames := make([]string, 0, len(services))
		for _, service := range services {
			serviceNames = append(serviceNames, service.Name)
		}
		mux.Handle(grpcreflect.NewHandlerV1(grpcreflect.NewStaticReflector(serviceNames...)))
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://0.0.0.0:5173",
		},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"Connect-Protocol-Version",
		},
	})

	handler := c.Handler(h2c.NewHandler(mux, &http2.Server{}))

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	slog.Info(
		"🚀 API server started successfully",
		slog.String("addr", addr),
		slog.String("env", string(cfg.Env)),
	)

	return http.ListenAndServe(addr, handler)
}
