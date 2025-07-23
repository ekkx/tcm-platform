package user

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/handler"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/service"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/usecase"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1/userv1connect"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool) userv1connect.UserServiceHandler {
	querier := database.New(dbPool)
	userRepo := repository.New(querier)
	userService := service.New(userRepo)
	userUseCase := usecase.New(userRepo, userService)
	return handler.New(userUseCase)
}
