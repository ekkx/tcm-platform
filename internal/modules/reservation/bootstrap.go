package reservation

import (
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1/reservationv1connect"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	rsvad "github.com/ekkx/tcm-platform/internal/modules/reservation/adapter"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/handler"
	rsvrepo "github.com/ekkx/tcm-platform/internal/modules/reservation/repository"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/usecase"
	subrepo "github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	userad "github.com/ekkx/tcm-platform/internal/modules/user/adapter"
	userrepo "github.com/ekkx/tcm-platform/internal/modules/user/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitModule(dbPool *pgxpool.Pool) reservationv1connect.ReservationServiceHandler {
	querier := sqlc.New(dbPool)
	userRepo := userrepo.New(querier)
	reservationRepo := rsvrepo.New(querier)
	subRepo := subrepo.New(querier)
	userQuery := userad.NewQueryAdapter(userRepo)
	reservationQuery := rsvad.NewQueryAdapter(reservationRepo)
	userAsm := assemble.NewUserAssembler(userQuery)
	reservationAsm := assemble.NewReservationAssembler(reservationQuery, userAsm)
	reservationUseCase := usecase.New(reservationRepo, subRepo, reservationAsm, userQuery)
	return handler.New(reservationUseCase)
}
