package server

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/server/interceptor"
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

type ServiceDefinition struct {
	Name            string
	RegisterHandler func(mux *http.ServeMux)
}

func getServiceDefinitions(cfg *config.Config, dbPool *pgxpool.Pool, jwtManager *jwt.JWTManager) []ServiceDefinition {
	return []ServiceDefinition{
		{
			Name: authv1connect.AuthServiceName,
			RegisterHandler: func(mux *http.ServeMux) {
				mux.Handle(authv1connect.NewAuthServiceHandler(
					auth.InitModule(dbPool, jwtManager),
					connect.WithInterceptors(
						interceptor.NewConfigInterceptor(cfg),
						interceptor.ErrorInterceptor(cfg.Env),
						interceptor.NewLoggingInterceptor(),
					),
				))
			},
		},
		{
			Name: reservationv1connect.ReservationServiceName,
			RegisterHandler: func(mux *http.ServeMux) {
				mux.Handle(reservationv1connect.NewReservationServiceHandler(
					reservation.InitModule(dbPool),
					connect.WithInterceptors(
						interceptor.NewConfigInterceptor(cfg),
						interceptor.ErrorInterceptor(cfg.Env),
						interceptor.NewLoggingInterceptor(),
						interceptor.AuthInterceptor(jwtManager),
						interceptor.UserVerificationInterceptor(dbPool),
					),
				))
			},
		},
		{
			Name: roomv1connect.RoomServiceName,
			RegisterHandler: func(mux *http.ServeMux) {
				mux.Handle(roomv1connect.NewRoomServiceHandler(
					room.InitModule(dbPool),
					connect.WithInterceptors(
						interceptor.NewConfigInterceptor(cfg),
						interceptor.ErrorInterceptor(cfg.Env),
						interceptor.NewLoggingInterceptor(),
						interceptor.AuthInterceptor(jwtManager),
						interceptor.UserVerificationInterceptor(dbPool),
					),
				))
			},
		},
		{
			Name: userv1connect.UserServiceName,
			RegisterHandler: func(mux *http.ServeMux) {
				mux.Handle(userv1connect.NewUserServiceHandler(
					user.InitModule(dbPool),
					connect.WithInterceptors(
						interceptor.NewConfigInterceptor(cfg),
						interceptor.ErrorInterceptor(cfg.Env),
						interceptor.NewLoggingInterceptor(),
						interceptor.AuthInterceptor(jwtManager),
						interceptor.UserVerificationInterceptor(dbPool),
					),
				))
			},
		},
	}
}
