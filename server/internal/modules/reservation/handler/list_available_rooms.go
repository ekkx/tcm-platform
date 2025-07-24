package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
)

func (h *HandlerImpl) ListAvailableRooms(ctx context.Context, req *connect.Request[reservationv1.ListAvailableRoomsRequest]) (*connect.Response[reservationv1.ListAvailableRoomsResponse], error) {
	input, err := usecase.NewListAvailableRoomsInputFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	output, err := h.useCase.ListAvailableRooms(ctx, input)
	if err != nil {
		return nil, err
	}

	return output.ToResponse(), nil
}
