package auth

import (
	"github.com/ekkx/tcm-platform/internal/gen/pb/auth/v1/authv1connect"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/modules/auth/handler"
	"github.com/ekkx/tcm-platform/internal/modules/auth/usecase"
	"github.com/ekkx/tcm-platform/internal/modules/user/adapter"
	"github.com/ekkx/tcm-platform/internal/modules/user/repository"
	"github.com/ekkx/tcm-platform/internal/platform/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool, jwtManager *jwt.JWTManager) authv1connect.AuthServiceHandler {
	q := sqlc.New(dbPool)
	userRepo := repository.New(q)
	userQuery := adapter.NewQueryAdapter(userRepo)
	userCmd := adapter.NewCommandAdapter(userRepo)
	authUC := usecase.New(jwtManager, userQuery, userCmd)
	return handler.New(authUC)
}
