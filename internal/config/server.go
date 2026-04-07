package config

type ServerConfig struct {
	Host           string   `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	Port           int      `env:"SERVER_PORT" envDefault:"8080"`
	AllowedOrigins []string `env:"SERVER_ALLOWED_ORIGINS" envSeparator:"," envDefault:"http://localhost:5173"`
}
