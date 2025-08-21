package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	"github.com/ekkx/tcmrsv-web/internal/shared/gateway"
)

type UseCase interface {
	GetReservation(ctx context.Context, input *GetReservationInput) (*GetReservationOutput, error)
	ListReservations(ctx context.Context, input *ListReservationsInput) (*ListReservationsOutput, error)
	CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error)
	DeleteReservation(ctx context.Context, input *DeleteReservationInput) (*DeleteReservationOutput, error)
}

type UseCaseImpl struct {
	reservationRepo repository.Repository
	reservationAsm  assemble.ReservationAssembler
	userQuery       gateway.UserQuery
}

func New(
	reservationRepo repository.Repository,
	reservationAsm assemble.ReservationAssembler,
	userQuery gateway.UserQuery,
) UseCase {
	return &UseCaseImpl{
		reservationRepo: reservationRepo,
		reservationAsm:  reservationAsm,
		userQuery:       userQuery,
	}
}
