package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/config"
	"github.com/ekkx/tcmrsv-web/internal/shared/ctxhelper"
)

func NewConfigInterceptor(cfg *config.Config) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			newCtx := ctxhelper.WithConfig(ctx, cfg)
			return next(newCtx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
