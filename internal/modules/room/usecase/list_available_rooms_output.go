package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	roomv1 "github.com/ekkx/tcm-platform/internal/gen/pb/room/v1"
)

type ListAvailableRoomsOutput struct {
	Rooms []*entity.Room
}

func NewListAvailableRoomsOutput(rooms []*entity.Room) *ListAvailableRoomsOutput {
	return &ListAvailableRoomsOutput{
		Rooms: rooms,
	}
}

func (st *ListAvailableRoomsOutput) ToResponse() *connect.Response[roomv1.ListAvailableRoomsResponse] {
	return connect.NewResponse(&roomv1.ListAvailableRoomsResponse{
		Rooms: presenter.ToRoomList(st.Rooms),
	})
}
