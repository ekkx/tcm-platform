package handler

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/usecase"
)

func (h *HandlerImpl) DeleteReservation(ctx context.Context, req *connect.Request[reservationv1.DeleteReservationRequest]) (*connect.Response[reservationv1.DeleteReservationResponse], error) {
	input, err := usecase.NewDeleteReservationInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.DeleteReservation(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
