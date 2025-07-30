package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/modules/reservation/usecase"
	reservationv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/reservation/v1"
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
