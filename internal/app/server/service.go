package server

import (
	"net/http"

	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/gen/pb/auth/v1/authv1connect"
	"github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1/reservationv1connect"
	"github.com/ekkx/tcm-platform/internal/gen/pb/room/v1/roomv1connect"
	"github.com/ekkx/tcm-platform/internal/gen/pb/user/v1/userv1connect"
	"github.com/ekkx/tcm-platform/internal/modules/auth"
	"github.com/ekkx/tcm-platform/internal/modules/reservation"
	"github.com/ekkx/tcm-platform/internal/modules/room"
	"github.com/ekkx/tcm-platform/internal/modules/user"
	"github.com/ekkx/tcm-platform/internal/platform/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Name            string
	RegisterHandler func(mux *http.ServeMux)
}

func initServices(cfg *config.Config, dbPool *pgxpool.Pool, jwtManager *jwt.JWTManager) []Service {
	base := BaseHandlerOptions(cfg)
	authed := AuthedHandlerOptions(cfg, jwtManager, dbPool)

	return []Service{
		{
			Name: authv1connect.AuthServiceName,
			RegisterHandler: func(mux *http.ServeMux) {
				mux.Handle(authv1connect.NewAuthServiceHandler(
					auth.InitModule(dbPool, jwtManager),
					base...,
				))
			},
		},
		{
			Name: reservationv1connect.ReservationServiceName,
			RegisterHandler: func(mux *http.ServeMux) {
				mux.Handle(reservationv1connect.NewReservationServiceHandler(
					reservation.InitModule(dbPool),
					authed...,
				))
			},
		},
		{
			Name: roomv1connect.RoomServiceName,
			RegisterHandler: func(mux *http.ServeMux) {
				mux.Handle(roomv1connect.NewRoomServiceHandler(
					room.InitModule(dbPool),
					authed...,
				))
			},
		},
		{
			Name: userv1connect.UserServiceName,
			RegisterHandler: func(mux *http.ServeMux) {
				mux.Handle(userv1connect.NewUserServiceHandler(
					user.InitModule(dbPool),
					authed...,
				))
			},
		},
	}
}
