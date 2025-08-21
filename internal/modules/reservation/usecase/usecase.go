package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/gateway"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/repository"
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
