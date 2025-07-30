package config

type LoggingConfig struct {
	Level string `env:"LOGGING_LEVEL" envDefault:"info"`
}
