package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/presenter"
)

type ListAvailableRoomsOutput struct {
	Rooms []*entity.Room
}

func NewListAvailableRoomsOutput(rooms []*entity.Room) *ListAvailableRoomsOutput {
	return &ListAvailableRoomsOutput{
		Rooms: rooms,
	}
}

func (st *ListAvailableRoomsOutput) ToResponse() *connect.Response[reservationv1.ListAvailableRoomsResponse] {
	return connect.NewResponse(&reservationv1.ListAvailableRoomsResponse{
		Rooms: presenter.ToRoomList(st.Rooms),
	})
}
