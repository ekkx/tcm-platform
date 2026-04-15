package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/gateway"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/repository"
	subrepo "github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
)

type UseCase interface {
	GetReservation(ctx context.Context, input *GetReservationInput) (*GetReservationOutput, error)
	ListReservations(ctx context.Context, input *ListReservationsInput) (*ListReservationsOutput, error)
	CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error)
	UpdateReservationNote(ctx context.Context, input *UpdateReservationNoteInput) (*UpdateReservationNoteOutput, error)
	DeleteReservation(ctx context.Context, input *DeleteReservationInput) (*DeleteReservationOutput, error)
}

type UseCaseImpl struct {
	reservationRepo repository.Repository
	subRepo         subrepo.Repository
	reservationAsm  assemble.ReservationAssembler
	userQuery       gateway.UserQuery
}

func New(
	reservationRepo repository.Repository,
	subRepo subrepo.Repository,
	reservationAsm assemble.ReservationAssembler,
	userQuery gateway.UserQuery,
) UseCase {
	return &UseCaseImpl{
		reservationRepo: reservationRepo,
		subRepo:         subRepo,
		reservationAsm:  reservationAsm,
		userQuery:       userQuery,
	}
}
