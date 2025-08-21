package room

import (
	"github.com/ekkx/tcm-platform/internal/gen/pb/room/v1/roomv1connect"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/modules/room/handler"
	roomrepo "github.com/ekkx/tcm-platform/internal/modules/room/repository"
	"github.com/ekkx/tcm-platform/internal/modules/room/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool) roomv1connect.RoomServiceClient {
	querier := sqlc.New(dbPool)
	roomRepo := roomrepo.New(querier)
	roomUseCase := usecase.New(roomRepo)
	return handler.New(roomUseCase)
}
