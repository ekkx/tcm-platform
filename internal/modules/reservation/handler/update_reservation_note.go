package handler

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/usecase"
)

func (h *HandlerImpl) UpdateReservationNote(ctx context.Context, req *connect.Request[reservationv1.UpdateReservationNoteRequest]) (*connect.Response[reservationv1.UpdateReservationNoteResponse], error) {
	input, err := usecase.NewUpdateReservationNoteInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.UpdateReservationNote(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
