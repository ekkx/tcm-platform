package commands

import (
	"github.com/ekkx/tcmrsv"
	authorization_hdl "github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/handler"
	authorization_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/usecase"
	reservation_hdl "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/handler"
	reservation_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	reservation_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	room_hdl "github.com/ekkx/tcmrsv-web/server/internal/modules/room/handler"
	room_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/room/usecase"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	authorization_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/authorization"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/room"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServerDeps struct {
	AuthorizationServiceServer authorization_v1.AuthorizationServiceServer
	ReservationServiceServer   reservation_v1.ReservationServiceServer
	RoomServiceServer          room_v1.RoomServiceServer
}

func GenerateServerDeps(pool *pgxpool.Pool) *ServerDeps {
	return &ServerDeps{
		AuthorizationServiceServer: authorization_hdl.NewHandler(authorization_uc.NewUsecase(tcmrsv.New(), user_repo.NewRepository(pool))),
		ReservationServiceServer:   reservation_hdl.NewHandler(reservation_uc.NewUsecase(reservation_repo.NewRepository(pool))),
		RoomServiceServer:          room_hdl.NewHandler(room_uc.NewUsecase()),
	}
}
