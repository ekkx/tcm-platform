package server

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/auth"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/interceptor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/auth/v1/authv1connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1/reservationv1connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1/userv1connect"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwt"
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
