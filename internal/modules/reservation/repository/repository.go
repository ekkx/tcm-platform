package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
)

type Repository interface {
	GetReservationByID(ctx context.Context, reservationID ulid.ULID) (*entity.Reservation, error)
	IsReservationConflicted(ctx context.Context, params *IsReservationConflictedParams) (bool, error)
	ListReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error)
	ListReservationIDsByDate(ctx context.Context, date ymd.YMD) ([]ulid.ULID, error)
	ListPendingReservationIDsByDate(ctx context.Context, date ymd.YMD) ([]ulid.ULID, error)
	ListUserReservationIDs(ctx context.Context, params *ListUserReservationIDsParams) ([]ulid.ULID, error)
	CreateReservation(ctx context.Context, params *CreateReservationParams) (*ulid.ULID, error)
	UpdateReservationStatus(ctx context.Context, params *UpdateReservationStatusParams) error
	UpdateReservationNote(ctx context.Context, params *UpdateReservationNoteParams) error
	DeleteReservationByID(ctx context.Context, reservationID ulid.ULID) error
}

type RepositoryImpl struct {
	querier sqlc.Querier
}

func New(querier sqlc.Querier) Repository {
	return &RepositoryImpl{
		querier: querier,
	}
}
