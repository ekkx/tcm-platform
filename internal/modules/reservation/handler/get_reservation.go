package handler

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/usecase"
)

func (h *HandlerImpl) GetReservation(ctx context.Context, req *connect.Request[reservationv1.GetReservationRequest]) (*connect.Response[reservationv1.GetReservationResponse], error) {
	input, err := usecase.NewGetReservationInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.GetReservation(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
