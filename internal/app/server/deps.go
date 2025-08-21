package server

import (
	"context"
	"fmt"

	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/platform/jwt"
	"github.com/ekkx/tcm-platform/internal/platform/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Deps struct {
	Cfg      *config.Config
	Pool     *pgxpool.Pool
	JWT      *jwt.JWTManager
	Services []Service
}

func initDependencies() (*Deps, func(), error) {
	cfg, err := config.New()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	logger.Init(cfg)

	ctx := context.Background()
	db, err := cfg.Database.Open(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	jwtMgr := jwt.NewJWTManager(cfg.Auth.JWTSecretKey, cfg.Auth.AccessExpiration, cfg.Auth.RefreshExpiration)

	svcs := initServices(cfg, db, jwtMgr)

	cleanup := func() { db.Close() }
	return &Deps{
		Cfg:      cfg,
		Pool:     db,
		JWT:      jwtMgr,
		Services: svcs,
	}, cleanup, nil
}
