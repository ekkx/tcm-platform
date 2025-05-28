package ctxhelper

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
)

type configKey struct{}
type accessTokenKey struct{}
type actorKey struct{}

func SetConfig(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, configKey{}, cfg)
}

func GetConfig(ctx context.Context) *config.Config {
	cfg, _ := ctx.Value(configKey{}).(*config.Config)
	return cfg
}

func SetAccessToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, accessTokenKey{}, token)
}

func GetAccessToken(ctx context.Context) string {
	token, _ := ctx.Value(accessTokenKey{}).(string)
	return token
}

func SetActor(ctx context.Context, actor actor.Actor) context.Context {
	return context.WithValue(ctx, actorKey{}, actor)
}

func GetActor(ctx context.Context) actor.Actor {
	token, _ := ctx.Value(actorKey{}).(actor.Actor)
	return token
}
