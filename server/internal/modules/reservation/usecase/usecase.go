package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	rsv_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	room_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
)

type Usecase interface {
	CreateReservation(ctx context.Context, params *input.CreateReservation) (*output.CreateReservation, error)
	GetReservation(ctx context.Context, params *input.GetReservation) (*output.GetReservation, error)
	GetUserReservations(ctx context.Context, params *input.GetUserReservations) (*output.GetMyReservations, error)
	UpdateReservation(ctx context.Context, params *input.UpdateReservation) (*output.UpdateReservation, error)
	DeleteReservation(ctx context.Context, params *input.DeleteReservation) error
}

type UsecaseImpl struct {
	rsvRepo  *rsv_repo.Repository
	roomRepo *room_repo.Repository
}

func NewUsecase(
	rsvRepo *rsv_repo.Repository,
	roomRepo *room_repo.Repository,
) *UsecaseImpl {
	return &UsecaseImpl{
		rsvRepo:  rsvRepo,
		roomRepo: roomRepo,
	}
}
