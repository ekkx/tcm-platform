package handler_test

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
)

// testRoomHandler is a test handler implementation
type testRoomHandler struct {
	room.UnimplementedRoomServiceServer
	mockGetRooms func(context.Context) *output.GetRooms
}

func (h *testRoomHandler) GetRooms(ctx context.Context, req *room.GetRoomsRequest) (*room.GetRoomsReply, error) {
	if h.mockGetRooms != nil {
		rooms := h.mockGetRooms(ctx)
		return rooms.ToProto(), nil
	}
	return &room.GetRoomsReply{
		Rooms: []*room.Room{},
	}, nil
}
