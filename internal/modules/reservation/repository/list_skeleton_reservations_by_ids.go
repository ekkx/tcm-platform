package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/internal/shared/mapper"
	"github.com/ekkx/tcmrsv-web/internal/shared/util"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

func (repo *RepositoryImpl) ListSkeletonReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error) {
	if len(reservationIDs) == 0 {
		return nil, nil
	}

	dbReservations, err := repo.querier.ListReservationsByIDs(ctx, util.ToULIDStrings(reservationIDs))
	if err != nil {
		return nil, err
	}

	reservations := make([]*entity.Reservation, 0, len(dbReservations))
	for _, r := range dbReservations {
		reservations = append(reservations, mapper.ToReservation(&r))
	}

	return reservations, nil
}
