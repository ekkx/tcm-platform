package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

type Repository interface {
	ListUnavailableRoomIDs(ctx context.Context, params *ListUnavailableRoomIDsParams) ([]string, error)
	ListSkeletonReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error)
	CreateReservation(ctx context.Context, params *CreateReservationParams) (*ulid.ULID, error)
}

type RepositoryImpl struct {
	querier database.Querier
}

func New(querier database.Querier) Repository {
	return &RepositoryImpl{
		querier: querier,
	}
}
