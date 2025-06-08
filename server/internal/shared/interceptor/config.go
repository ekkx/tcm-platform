package interceptor

import (
	"context"

	"google.golang.org/grpc"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
)

func ConfigUnaryInterceptor(cfg *config.Config) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = ctxhelper.SetConfig(ctx, cfg)
		return handler(ctx, req)
	}
}
