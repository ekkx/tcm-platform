package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/usecase"
	roomv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/room/v1"
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
