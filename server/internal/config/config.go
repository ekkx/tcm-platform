package config

import (
	"github.com/caarlos0/env/v11"
)

type Env string

const (
	EnvDevelopment Env = "development"
	EnvStaging     Env = "staging"
	EnvProduction  Env = "production"
)

type Config struct {
	Env      Env `env:"ENVIRONMENT"`
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
	Log      LoggingConfig
}

func New() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	if err := env.Parse(&cfg.Server); err != nil {
		return nil, err
	}
	if err := env.Parse(&cfg.Database); err != nil {
		return nil, err
	}
	if err := env.Parse(&cfg.Auth); err != nil {
		return nil, err
	}
	if err := env.Parse(&cfg.Log); err != nil {
		return nil, err
	}

	return cfg, nil
}
