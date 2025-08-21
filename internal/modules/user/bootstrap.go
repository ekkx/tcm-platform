package user

import (
	"github.com/ekkx/tcmrsv-web/internal/modules/user/adapter"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/handler"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/usecase"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	"github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1/userv1connect"
	"github.com/ekkx/tcmrsv-web/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool) userv1connect.UserServiceHandler {
	querier := database.New(dbPool)
	userRepo := repository.New(querier)
	userQuery := adapter.NewQueryAdapter(userRepo)
	userAsm := assemble.NewUserAssembler(userQuery)
	userUC := usecase.New(userRepo, userAsm)
	return handler.New(userUC)
}
