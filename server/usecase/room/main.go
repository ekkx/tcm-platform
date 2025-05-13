package room

import (
	"context"

	"github.com/ekkx/tcmrsv"
)

type RoomUsecase interface {
	GetRooms(ctx context.Context) *GetRoomsOutput
}

type RoomUsecaseImpl struct {
	tcmClient *tcmrsv.Client
}

var _ RoomUsecase = (*RoomUsecaseImpl)(nil)

func New(tcmClient *tcmrsv.Client) RoomUsecase {
	return &RoomUsecaseImpl{
		tcmClient: tcmClient,
	}
}
