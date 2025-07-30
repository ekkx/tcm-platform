package config

import "time"

type AuthConfig struct {
    JWTSecretKey      string        `env:"AUTH_JWT_SECRET_KEY"`
	AccessExpiration  time.Duration `env:"AUTH_ACCESS_EXPIRATION" envDefault:"24h"`
	RefreshExpiration time.Duration `env:"AUTH_REFRESH_EXPIRATION" envDefault:"168h"`
}
