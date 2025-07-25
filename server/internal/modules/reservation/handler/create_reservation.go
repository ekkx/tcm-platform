package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
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
