package commands

import (
	"github.com/ekkx/tcmrsv"
	auth_hdl "github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/handler"
	auth_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/usecase"
	rsv_hdl "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/handler"
	rsv_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	rsv_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	room_hdl "github.com/ekkx/tcmrsv-web/server/internal/modules/room/handler"
	room_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
	room_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/room/usecase"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/authorization"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServerDeps struct {
	AuthorizationServiceServer auth_v1.AuthorizationServiceServer
	ReservationServiceServer   rsv_v1.ReservationServiceServer
	RoomServiceServer          room_v1.RoomServiceServer
}

func GenerateServerDeps(pool *pgxpool.Pool) *ServerDeps {
	return &ServerDeps{
		AuthorizationServiceServer: auth_hdl.NewHandler(auth_uc.NewUsecase(tcmrsv.New(), user_repo.NewRepository(pool))),
		ReservationServiceServer:   rsv_hdl.NewHandler(rsv_uc.NewUsecase(rsv_repo.NewRepository(pool))),
		RoomServiceServer:          room_hdl.NewHandler(room_uc.NewUsecase(room_repo.NewRepository(tcmrsv.New()))),
	}
}
