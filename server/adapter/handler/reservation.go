package handler

import (
	"net/http"

	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/pkg/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/response"
	"github.com/ekkx/tcmrsv-web/server/usecase/reservation"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetReservation(ctx echo.Context, id string) error {
	return nil
}

func (h *Handler) GetMyReservations(ctx echo.Context) error {
	stdCtx := ctx.Request().Context()

	output, err := h.reservationUsecase.GetMyReservations(stdCtx, &reservation.GetMyReservationsInput{
		UserID:   ctxhelper.GetRequestUser(stdCtx).ID,
		Password: ctxhelper.GetRequestUser(stdCtx).Password,
	})
	if err != nil {
		return err
	}

	var apiReservations []api.Reservation
	for _, rsv := range output.Reservations {
		apiReservations = append(apiReservations, api.Reservation{
			Id:     rsv.ID,
			RoomId: rsv.RoomID,
		})
	}

	return response.JSON(ctx, http.StatusOK, apiReservations)
}

func (h *Handler) CreateReservation(ctx echo.Context) error {
	return nil
}

func (h *Handler) UpdateReservation(ctx echo.Context, id string) error {
	return nil
}

func (h *Handler) DeleteReservation(ctx echo.Context, id string) error {
	return nil
}
