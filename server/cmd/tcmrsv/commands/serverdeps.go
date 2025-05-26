package commands

import (
	"github.com/ekkx/tcmrsv"
	authorize_hdl "github.com/ekkx/tcmrsv-web/server/internal/modules/authorize/handler"
	authorize_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/authorize/usecase"
	reservation_hdl "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/handler"
	reservation_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	reservation_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	room_hdl "github.com/ekkx/tcmrsv-web/server/internal/modules/room/handler"
	room_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/room/usecase"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	authorize_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/authorize"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/room"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServerDeps struct {
	AuthorizeServiceServer   authorize_v1.AuthorizeServiceServer
	ReservationServiceServer reservation_v1.ReservationServiceServer
	RoomServiceServer        room_v1.RoomServiceServer
}

func GenerateServerDeps(pool *pgxpool.Pool) *ServerDeps {
	return &ServerDeps{
		AuthorizeServiceServer:   authorize_hdl.NewHandler(authorize_uc.NewUsecase(user_repo.NewRepository(pool), tcmrsv.New())),
		ReservationServiceServer: reservation_hdl.NewHandler(reservation_uc.NewUsecase(reservation_repo.NewRepository(pool))),
		RoomServiceServer:        room_hdl.NewHandler(room_uc.NewUsecase()),
	}
}
