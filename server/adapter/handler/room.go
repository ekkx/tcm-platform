package handler

import (
	"net/http"

	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/pkg/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRooms(ctx echo.Context, params api.GetRoomsParams) error {
	return response.JSON(ctx, http.StatusOK, struct{}{})
}
