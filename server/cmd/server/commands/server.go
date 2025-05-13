package commands

import (
	"log"
	"net/http"

	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/adapter/handler"
	"github.com/ekkx/tcmrsv-web/server/adapter/middleware"
	"github.com/ekkx/tcmrsv-web/server/config"
	"github.com/ekkx/tcmrsv-web/server/infra/db"
	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
	"github.com/ekkx/tcmrsv-web/server/usecase/authorize"
	"github.com/ekkx/tcmrsv-web/server/usecase/reservation"
	"github.com/ekkx/tcmrsv-web/server/usecase/room"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/ekkx/tcmrsv"

	_ "time/tzdata"
)

func Run(cfg *config.Config) {
	pool, err := db.Open(&db.ConnectionOptions{
		ConnString:      cfg.DBDSN,
		MaxConnLifetime: cfg.DBMaxConnLifetime,
		MaxConnIdleTime: cfg.DBMaxConnIdleTime,
		MaxConns:        cfg.DBMaxConns,
		MinConns:        cfg.DBMinConns,
	})
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer pool.Close()

	e := echo.New()

	e.HTTPErrorHandler = apperrors.ErrorHandler

	middlewarecfg := echomiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.CORSWithConfig(middlewarecfg))
	e.Use(middleware.Config(cfg))
	e.Use(middleware.JWT(pool, []string{"/authorize", "/authorize/refresh"}))
	e.Use(middleware.RequestValidator(swagger))

	querier := db.New(pool)
	tcmClient := tcmrsv.New()

	h := handler.New(
		authorize.New(tcmClient, querier),
		reservation.New(tcmClient, querier),
		room.New(tcmClient),
	)

	api.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(":1323"))
}
