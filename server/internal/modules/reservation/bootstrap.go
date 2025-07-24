package reservation

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/handler"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1/reservationv1connect"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool) reservationv1connect.ReservationServiceHandler {
	querier := database.New(dbPool)
	reservationRepo := repository.New(querier)
	reservationUseCase := usecase.New(reservationRepo)
	return handler.New(reservationUseCase)
}
