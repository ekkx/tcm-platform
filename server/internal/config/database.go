package config

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConfig struct {
	DSN             string        `env:"DATABASE_DSN"`
	MaxConns        int32         `env:"DATABASE_MAX_CONNS"`
	MinConns        int32         `env:"DATABASE_MIN_CONNS"`
	MaxConnLifetime time.Duration `env:"DATABASE_MAX_CONN_LIFETIME"`
	MaxConnIdleTime time.Duration `env:"DATABASE_MAX_CONN_IDLE_TIME"`
}

func (cfg *DatabaseConfig) Open(ctx context.Context) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(
		cfg.DSN,
	)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = cfg.MaxConns
	poolCfg.MinConns = cfg.MinConns
	poolCfg.MaxConnLifetime = cfg.MaxConnLifetime
	poolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime

	return pgxpool.NewWithConfig(ctx, poolCfg)
}
