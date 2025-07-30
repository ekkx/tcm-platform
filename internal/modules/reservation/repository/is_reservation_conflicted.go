package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/pkg/database"
	"github.com/ekkx/tcmrsv-web/pkg/ymd"
)

type IsReservationConflictedParams struct {
	RoomID     string
	Date       ymd.YMD
	FromHour   int
	FromMinute int
	ToHour     int
	ToMinute   int
}

func (repo *RepositoryImpl) IsReservationConflicted(ctx context.Context, params *IsReservationConflictedParams) (bool, error) {
	fromHour := int32(params.FromHour)
	fromMinute := int32(params.FromMinute)
	toHour := int32(params.ToHour)
	toMinute := int32(params.ToMinute)

	isConflicted, err := repo.querier.IsReservationConflicted(ctx, database.IsReservationConflictedParams{
		RoomID:     params.RoomID,
		Date:       params.Date,
		FromHour:   &fromHour,
		FromMinute: &fromMinute,
		ToHour:     &toHour,
		ToMinute:   &toMinute,
	})
	if err != nil {
		return false, err
	}

	return isConflicted, nil
}
