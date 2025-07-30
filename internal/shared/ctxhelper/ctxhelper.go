package ctxhelper

import (
	"context"
	"log/slog"

	"github.com/ekkx/tcmrsv-web/internal/config"
	"github.com/ekkx/tcmrsv-web/pkg/actor"
)

type configKey struct{}

func WithConfig(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, configKey{}, cfg)
}

func Config(ctx context.Context) *config.Config {
	cfg, _ := ctx.Value(configKey{}).(*config.Config)
	return cfg
}

type loggerKey struct{}

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func Logger(ctx context.Context) *slog.Logger {
	logger, _ := ctx.Value(loggerKey{}).(*slog.Logger)
	return logger
}

type actorKey struct{}

func WithActor(ctx context.Context, actor *actor.Actor) context.Context {
	return context.WithValue(ctx, actorKey{}, actor)
}

func Actor(ctx context.Context) *actor.Actor {
	actor, _ := ctx.Value(actorKey{}).(*actor.Actor)
	return actor
}
