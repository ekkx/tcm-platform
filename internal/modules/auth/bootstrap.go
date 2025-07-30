package auth

import (
	"github.com/ekkx/tcmrsv-web/internal/modules/auth/handler"
	"github.com/ekkx/tcmrsv-web/internal/modules/auth/usecase"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/service"
	"github.com/ekkx/tcmrsv-web/internal/shared/pb/auth/v1/authv1connect"
	"github.com/ekkx/tcmrsv-web/pkg/database"
	"github.com/ekkx/tcmrsv-web/pkg/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool, jwtManager *jwt.JWTManager) authv1connect.AuthServiceHandler {
	querier := database.New(dbPool)
	userRepository := repository.New(querier)
	userService := service.New(userRepository)
	authUseCase := usecase.New(jwtManager, userRepository, userService)
	return handler.New(authUseCase)
}
