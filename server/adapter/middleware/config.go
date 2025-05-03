package middleware

import (
	"github.com/ekkx/tcmrsv-web/server/config"
	"github.com/ekkx/tcmrsv-web/server/pkg/ctxhelper"
	"github.com/labstack/echo/v4"
)

func Config(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c := ctxhelper.SetConfig(ctx.Request().Context(), cfg)
			ctx.SetRequest(ctx.Request().WithContext(c))
			return next(ctx)
		}
	}
}
