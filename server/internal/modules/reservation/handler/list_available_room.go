package handler

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
)

func (h *Handler) ListAvailableRooms(ctx context.Context, req *connect.Request[reservationv1.ListAvailableRoomsRequest]) (*connect.Response[reservationv1.ListAvailableRoomsResponse], error) {
    return nil, nil
}
