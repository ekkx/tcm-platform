package commands

import (
	"fmt"
	"net/http"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/auth"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/auth/v1/authv1connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1/reservationv1connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1/userv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Run(cfg *config.Config) error {
    mux := http.NewServeMux()

    mux.Handle(authv1connect.NewAuthServiceHandler(auth.InitModule()))
    mux.Handle(reservationv1connect.NewReservationServiceHandler(reservation.InitModule()))
    mux.Handle(userv1connect.NewUserServiceHandler(user.InitModule()))

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:5173"},
        AllowCredentials: true,
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders: []string{
            "Content-Type",
            "Authorization",
            "Connect-Protocol-Version",
        },
        // Enable Debugging for testing, consider disabling in production
        Debug: true,
    })

    handler := c.Handler(h2c.NewHandler(mux, &http2.Server{}))

    http.ListenAndServe(fmt.Sprintf(":%d", 50051), handler)

    return nil
}
