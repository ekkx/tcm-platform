package handler

import (
	"net/http"

	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/pkg/response"
	"github.com/ekkx/tcmrsv-web/server/usecase/authorize"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Authorize(ctx echo.Context) error {
	var body api.AuthorizeJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	output, err := h.authorizeUsecase.Login(ctx.Request().Context(), &authorize.LoginInput{
		UserID:   body.UserId,
		Password: body.Password,
	})
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, &api.Authorization{
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	})
}
