package handler

import (
	"net/http"

	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/adapter/presenter"
	"github.com/ekkx/tcmrsv-web/server/domain"
	"github.com/ekkx/tcmrsv-web/server/pkg/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/response"
	"github.com/ekkx/tcmrsv-web/server/usecase/reservation"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetReservation(ctx echo.Context, id string) error {
	return nil
}

func (h *Handler) GetMyReservations(ctx echo.Context, params api.GetMyReservationsParams) error {
	stdCtx := ctx.Request().Context()

	output, err := h.reservationUsecase.GetMyReservations(stdCtx, &reservation.GetMyReservationsInput{
		UserID:    ctxhelper.GetRequestUser(stdCtx).ID,
		Password:  ctxhelper.GetRequestUser(stdCtx).Password,
		Campus:    (*domain.Campus)(params.CampusCode),
		PianoType: (*domain.PianoType)(params.PianoType),
		Date:      params.Date,
	})
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, presenter.GetMyReservations(output))
}

func (h *Handler) CreateReservation(ctx echo.Context) error {
	var body api.CreateReservationJSONBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	stdCtx := ctx.Request().Context()

	var pianoTypes *[]domain.PianoType
	if body.PianoTypes != nil {
		pts := make([]domain.PianoType, len(*body.PianoTypes))
		for i, pt := range *body.PianoTypes {
			pts[i] = domain.PianoType(pt)
		}
		pianoTypes = &pts
	}

	output, err := h.reservationUsecase.CreateReservation(stdCtx, &reservation.CreateReservationInput{
		UserID:       ctxhelper.GetRequestUser(stdCtx).ID,
		Password:     ctxhelper.GetRequestUser(stdCtx).Password,
		Campus:       domain.Campus(body.CampusCode),
		Date:         body.Date,
		FromHour:     body.FromHour,
		FromMinute:   body.FromMinute,
		ToHour:       body.ToHour,
		ToMinute:     body.ToMinute,
		IsAutoSelect: body.IsAutoSelect,
		RoomID:       body.RoomId,
		BookerName:   body.BookerName,
		PianoNumbers: body.PianoNumbers,
		PianoTypes:   pianoTypes,
		Floors:       body.Floors,
		IsBasement:   body.IsBasement,
	})
	if err != nil {
		return err
	}

	return response.JSON(ctx, http.StatusOK, presenter.CreateReservation(output))
}

func (h *Handler) UpdateReservation(ctx echo.Context, id string) error {
	return nil
}

func (h *Handler) DeleteReservation(ctx echo.Context, id string) error {
	return nil
}
