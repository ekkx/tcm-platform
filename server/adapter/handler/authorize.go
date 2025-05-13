package handler

import (
	"net/http"

	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/pkg/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/response"
	"github.com/ekkx/tcmrsv-web/server/usecase/authorize"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Authorize(ctx echo.Context) error {
	var body api.AuthorizeJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	stdCtx := ctx.Request().Context()

	output, err := h.authorizeUsecase.Authorize(stdCtx, &authorize.AuthorizeInput{
		UserID:         body.UserId,
		Password:       body.Password,
		PasswordAESKey: ctxhelper.GetConfig(stdCtx).PasswordAESKey,
		JWTSecret:      ctxhelper.GetConfig(stdCtx).JWTSecret,
	})
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, &api.Authorization{
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	})
}

func (h *Handler) Reauthorize(ctx echo.Context) error {
	var body api.ReauthorizeJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	stdCtx := ctx.Request().Context()

	output, err := h.authorizeUsecase.Reauthorize(stdCtx, &authorize.ReauthorizeInput{
		RefreshToken:   body.RefreshToken,
		PasswordAESKey: ctxhelper.GetConfig(stdCtx).PasswordAESKey,
		JWTSecret:      ctxhelper.GetConfig(stdCtx).JWTSecret,
	})
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, &api.Authorization{
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	})
}
