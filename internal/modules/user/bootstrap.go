package user

import (
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/gen/pb/user/v1/userv1connect"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/modules/user/adapter"
	"github.com/ekkx/tcm-platform/internal/modules/user/handler"
	"github.com/ekkx/tcm-platform/internal/modules/user/repository"
	"github.com/ekkx/tcm-platform/internal/modules/user/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool) userv1connect.UserServiceHandler {
	querier := sqlc.New(dbPool)
	userRepo := repository.New(querier)
	userQuery := adapter.NewQueryAdapter(userRepo)
	userAsm := assemble.NewUserAssembler(userQuery)
	userUC := usecase.New(userRepo, userAsm)
	return handler.New(userUC)
}
