package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

func (repo *RepositoryImpl) GetReservationByID(ctx context.Context, reservationID ulid.ULID) (*entity.Reservation, error) {
	reservations, err := repo.ListReservationsByIDs(ctx, []ulid.ULID{reservationID})
	if err != nil {
		return nil, err
	}
	if len(reservations) == 0 {
		return nil, nil
	}
	return reservations[0], nil
}
