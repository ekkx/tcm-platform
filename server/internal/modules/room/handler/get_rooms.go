package handler

import (
	"context"

	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/room"
)

func (h *Handler) GetRooms(ctx context.Context, req *room_v1.GetRoomsRequest) (*room_v1.GetRoomsReply, error) {
	rooms := h.Usecase.GetRooms(ctx)
	return rooms.ToProto(), nil
}
