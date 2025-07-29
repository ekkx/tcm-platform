package service

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
	"github.com/ekkx/tcmrsv-web/server/pkg/ymd"
)

type ListUserReservationsParams struct {
	UserID ulid.ULID
	Date   ymd.YMD
}

func (svc *ServiceImpl) ListUserReservations(ctx context.Context, params *ListUserReservationsParams) ([]*entity.Reservation, error) {
	rsvIDs, err := svc.reservationRepo.ListUserReservationIDs(ctx, &repository.ListUserReservationIDsParams{
		UserID: params.UserID,
		Date:   params.Date,
	})
	if err != nil {
		return nil, err
	}

	if len(rsvIDs) == 0 {
		return nil, nil
	}

	return svc.ListReservationsByIDs(ctx, rsvIDs)
}
