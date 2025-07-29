package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/service"
)

type UseCase interface {
	GetReservation(ctx context.Context, input *GetReservationInput) (*GetReservationOutput, error)
	ListReservations(ctx context.Context, input *ListReservationsInput) (*ListReservationsOutput, error)
	CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error)
	DeleteReservation(ctx context.Context, input *DeleteReservationInput) (*DeleteReservationOutput, error)
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
