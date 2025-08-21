package handler

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/usecase"
)

func (h *HandlerImpl) CreateReservation(ctx context.Context, req *connect.Request[reservationv1.CreateReservationRequest]) (*connect.Response[reservationv1.CreateReservationResponse], error) {
	input, err := usecase.NewCreateReservationInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.CreateReservation(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
