package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

type Repository interface {
	IsReservationConflicted(ctx context.Context, params *IsReservationConflictedParams) (bool, error)
	ListSkeletonReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error)
	ListUserReservationIDs(ctx context.Context, params *ListUserReservationIDsParams) ([]ulid.ULID, error)
	CreateReservation(ctx context.Context, params *CreateReservationParams) (*ulid.ULID, error)
	DeleteReservationByID(ctx context.Context, reservationID ulid.ULID) error
}

type RepositoryImpl struct {
	querier database.Querier
}

func New(querier database.Querier) Repository {
	return &RepositoryImpl{
		querier: querier,
	}
}
