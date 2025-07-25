package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
	"github.com/ekkx/tcmrsv-web/server/pkg/ymd"
)

type CreateReservationParams struct {
	ID         ulid.ULID
	UserID     ulid.ULID
	CampusType enum.CampusType
	RoomID     string
	Date       ymd.YMD
	FromHour   int
	FromMinute int
	ToHour     int
	ToMinute   int
}

func (repo *RepositoryImpl) CreateReservation(ctx context.Context, params *CreateReservationParams) (*ulid.ULID, error) {
	if params.ID.IsZero() {
		params.ID = ulid.New()
	}

	var campusType database.CampusType
	switch params.CampusType {
	case enum.CampusTypeIkebukuro:
		campusType = database.CampusTypeIkebukuro
	case enum.CampusTypeNakameguro:
		campusType = database.CampusTypeNakameguro
	default:
		return nil, errs.ErrInvalidCampusType
	}

	id, err := repo.querier.CreateReservation(ctx, database.CreateReservationParams{
		ID:         params.ID,
		UserID:     params.UserID,
		CampusType: campusType,
		RoomID:     params.RoomID,
		Date:       params.Date,
		FromHour:   int32(params.FromHour),
		FromMinute: int32(params.FromMinute),
		ToHour:     int32(params.ToHour),
		ToMinute:   int32(params.ToMinute),
	})
	if err != nil {
		return nil, err
	}

	return &id, nil
}
