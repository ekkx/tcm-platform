package ctxhelper

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/pkg/config"
)

type configKey struct{}
type accessTokenKey struct{}
type requestUserKey struct{}

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

type RequestUser struct {
	ID       string
	Password string
}

func SetRequestUser(ctx context.Context, user RequestUser) context.Context {
	return context.WithValue(ctx, requestUserKey{}, user)
}

func GetRequestUser(ctx context.Context) RequestUser {
	token, _ := ctx.Value(requestUserKey{}).(RequestUser)
	return token
}
