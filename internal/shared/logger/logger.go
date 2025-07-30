package logger

import (
	"log/slog"
	"os"

	"github.com/ekkx/tcmrsv-web/internal/config"
	"github.com/lmittmann/tint"
)

func parseLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func Init(cfg *config.Config) {
	var handler slog.Handler

	switch cfg.Env {
	case config.EnvDevelopment:
		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level:      parseLevel(cfg.Log.Level),
			TimeFormat: "15:04:05.000",
		})
	default:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: parseLevel(cfg.Log.Level),
		})
	}

	slog.SetDefault(slog.New(handler))
}
