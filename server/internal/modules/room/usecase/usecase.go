package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
)

type UseCase interface {
	ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error)
}

type UseCaseImpl struct {
	roomRepo repository.Repository
}

func New(
	roomRepo repository.Repository,
) UseCase {
	return &UseCaseImpl{
		roomRepo: roomRepo,
	}
}
