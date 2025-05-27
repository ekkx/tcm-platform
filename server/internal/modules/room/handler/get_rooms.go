package handler

import (
	"context"

	room_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/room"
)

func (h *Handler) GetRooms(ctx context.Context, req *room_v1.GetRoomsRequest) (*room_v1.GetRoomsReply, error) {
	// TODO: Implement the logic to retrieve rooms.
	return nil, nil
}
