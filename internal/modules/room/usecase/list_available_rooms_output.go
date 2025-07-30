package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	roomv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/room/v1"
	"github.com/ekkx/tcmrsv-web/internal/shared/presenter"
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
