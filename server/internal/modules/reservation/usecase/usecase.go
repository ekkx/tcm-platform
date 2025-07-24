package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
)

type UseCase interface {
	ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error)
}

type UseCaseImpl struct {
	reservationRepo repository.Repository
}

func New(
	reservationRepo repository.Repository,
) UseCase {
	return &UseCaseImpl{
		reservationRepo: reservationRepo,
	}
}
