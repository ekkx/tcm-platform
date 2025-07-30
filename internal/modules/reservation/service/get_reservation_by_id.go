package service

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

func (svc *ServiceImpl) GetReservationByID(ctx context.Context, id ulid.ULID) (*entity.Reservation, error) {
	reservations, err := svc.ListReservationsByIDs(ctx, []ulid.ULID{id})
	if err != nil {
		return nil, err
	}
	if len(reservations) == 0 {
		return nil, nil
	}
	return reservations[0], nil
}
