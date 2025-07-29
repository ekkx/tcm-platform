package room

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/handler"
	roomrepo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/usecase"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/room/v1/roomv1connect"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool) roomv1connect.RoomServiceClient {
	querier := database.New(dbPool)
	roomRepo := roomrepo.New(querier)
	roomUseCase := usecase.New(roomRepo)
	return handler.New(roomUseCase)
}
