package handler

import (
	"context"

	"connectrpc.com/connect"
	roomv1 "github.com/ekkx/tcm-platform/internal/gen/pb/room/v1"
	"github.com/ekkx/tcm-platform/internal/modules/room/usecase"
)

func (h *HandlerImpl) ListAvailableRooms(ctx context.Context, req *connect.Request[roomv1.ListAvailableRoomsRequest]) (*connect.Response[roomv1.ListAvailableRoomsResponse], error) {
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
