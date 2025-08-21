package handler

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/usecase"
)

func (h *HandlerImpl) ListReservations(ctx context.Context, req *connect.Request[reservationv1.ListReservationsRequest]) (*connect.Response[reservationv1.ListReservationsResponse], error) {
	input, err := usecase.NewListReservationsInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.ListReservations(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
