package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
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
