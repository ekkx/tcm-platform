package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/mapper"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/util"
)

func (repo *RepositoryImpl) ListReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error) {
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
