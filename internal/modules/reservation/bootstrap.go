package reservation

import (
	rsvad "github.com/ekkx/tcmrsv-web/internal/modules/reservation/adapter"
	"github.com/ekkx/tcmrsv-web/internal/modules/reservation/handler"
	rsvrepo "github.com/ekkx/tcmrsv-web/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/internal/modules/reservation/usecase"
	userad "github.com/ekkx/tcmrsv-web/internal/modules/user/adapter"
	userrepo "github.com/ekkx/tcmrsv-web/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	"github.com/ekkx/tcmrsv-web/internal/shared/pb/reservation/v1/reservationv1connect"
	"github.com/ekkx/tcmrsv-web/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool) reservationv1connect.ReservationServiceHandler {
	querier := database.New(dbPool)
	userRepo := userrepo.New(querier)
	reservationRepo := rsvrepo.New(querier)
	userQuery := userad.NewQueryAdapter(userRepo)
	reservationQuery := rsvad.NewQueryAdapter(reservationRepo)
	userAsm := assemble.NewUserAssembler(userQuery)
	reservationAsm := assemble.NewReservationAssembler(reservationQuery, userAsm)
	reservationUseCase := usecase.New(reservationRepo, reservationAsm, userQuery)
	return handler.New(reservationUseCase)
}
