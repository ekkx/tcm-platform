package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/auth"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/interceptor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/logger"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/auth/v1/authv1connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1/reservationv1connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1/userv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	logger.Init(cfg)

	interceptors := connect.WithInterceptors(
		interceptor.NewConfigInterceptor(cfg),
		interceptor.NewLoggingInterceptor(),
	)

	mux := http.NewServeMux()

	mux.Handle(authv1connect.NewAuthServiceHandler(auth.InitModule(), interceptors))
	mux.Handle(reservationv1connect.NewReservationServiceHandler(reservation.InitModule(), interceptors))
	mux.Handle(userv1connect.NewUserServiceHandler(user.InitModule(), interceptors))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"Connect-Protocol-Version",
		},
	})

	handler := c.Handler(h2c.NewHandler(mux, &http2.Server{}))

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	slog.Info(
		"🚀 API server started successfully",
		slog.String("addr", addr),
		slog.String("env", string(cfg.Env)),
	)

	http.ListenAndServe(addr, handler)

	return nil
}
