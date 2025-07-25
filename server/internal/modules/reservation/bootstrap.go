package reservation

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/handler"
	rsvrepo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	rsvsvc "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/service"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	userrepo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	usersvc "github.com/ekkx/tcmrsv-web/server/internal/modules/user/service"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1/reservationv1connect"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool) reservationv1connect.ReservationServiceHandler {
	querier := database.New(dbPool)
	userRepo := userrepo.New(querier)
	reservationRepo := rsvrepo.New(querier)
	userService := usersvc.New(userRepo)
	reservationService := rsvsvc.New(reservationRepo, userService)
	reservationUseCase := usecase.New(reservationRepo, reservationService)
	return handler.New(reservationUseCase)
}
