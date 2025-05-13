package handler

import (
	"net/http"

	"github.com/ekkx/tcmrsv-web/server/adapter/presenter"
	"github.com/ekkx/tcmrsv-web/server/pkg/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRooms(ctx echo.Context) error {
	output := h.roomUsecase.GetRooms(ctx.Request().Context())
	return response.JSON(ctx, http.StatusOK, presenter.GetRooms(output))
}
