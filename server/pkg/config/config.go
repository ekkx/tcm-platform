package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	APPEnv string `env:"APP_ENV" envDefault:"development"`

	Database DatabaseConfig

	JWTSecret      string `env:"JWT_SECRET" envDefault:"jwt_secret"`
	PasswordAESKey string `env:"PASSWORD_AES_KEY" envDefault:"password_aes_key"`
}

func New() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	if err := env.Parse(&cfg.Database); err != nil {
		return nil, err
	}

	return cfg, nil
}
