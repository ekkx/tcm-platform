package server

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/server/interceptor"
	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/platform/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BaseHandlerOptions(cfg *config.Config) []connect.HandlerOption {
	return []connect.HandlerOption{
		connect.WithInterceptors(
			interceptor.NewConfigInterceptor(cfg),
			interceptor.ErrorInterceptor(cfg.Env),
			interceptor.NewLoggingInterceptor(),
		),
	}
}

func AuthedHandlerOptions(cfg *config.Config, jwtMgr *jwt.JWTManager, db *pgxpool.Pool) []connect.HandlerOption {
	base := BaseHandlerOptions(cfg)
	base[0] = connect.WithInterceptors(
		interceptor.NewConfigInterceptor(cfg),
		interceptor.ErrorInterceptor(cfg.Env),
		interceptor.NewLoggingInterceptor(),
		interceptor.AuthInterceptor(jwtMgr),
		interceptor.UserVerificationInterceptor(db), // TODO: recieve gateway.UserQuery
	)
	return base
}
