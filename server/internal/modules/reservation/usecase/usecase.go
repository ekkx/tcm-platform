package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/service"
)

type UseCase interface {
	ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error)
	CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error)
}

type UseCaseImpl struct {
	reservationRepo    repository.Repository
	reservationService service.Service
}

func New(
	reservationRepo repository.Repository,
	reservationService service.Service,
) UseCase {
	return &UseCaseImpl{
		reservationRepo:    reservationRepo,
		reservationService: reservationService,
	}
}
