package config

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConfig struct {
	DBDSN             string        `env:"DB_DSN" envDefault:"postgres://user:password@db:5432/db"`
	DBMaxConnLifetime time.Duration `env:"DB_MAX_CONN_LIFETIME" envDefault:"30m"`
	DBMaxConnIdleTime time.Duration `env:"DB_MAX_CONN_IDLE_TIME" envDefault:"5m"`
	DBMaxConns        int32         `env:"DB_MAX_CONNS" envDefault:"10"`
	DBMinConns        int32         `env:"DB_MIN_CONNS" envDefault:"1"`
}

func (cfg *DatabaseConfig) Open() (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(
		cfg.DBDSN,
	)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConnLifetime = cfg.DBMaxConnLifetime
	poolCfg.MaxConnIdleTime = cfg.DBMaxConnIdleTime
	poolCfg.MaxConns = cfg.DBMaxConns
	poolCfg.MinConns = cfg.DBMinConns

	return pgxpool.NewWithConfig(context.Background(), poolCfg)
}
