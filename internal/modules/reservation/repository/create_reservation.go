package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
)

type CreateReservationParams struct {
	ID         ulid.ULID
	UserID     ulid.ULID
	CampusType valueobject.CampusType
	RoomID     string
	Date       ymd.YMD
	FromHour   int
	FromMinute int
	ToHour     int
	ToMinute   int
	Note       *string
}

func (repo *RepositoryImpl) CreateReservation(ctx context.Context, params *CreateReservationParams) (*ulid.ULID, error) {
	if params.ID.IsZero() {
		params.ID = ulid.New()
	}

	var campusType sqlc.CampusType
	switch params.CampusType {
	case valueobject.CampusTypeIkebukuro:
		campusType = sqlc.CampusTypeIkebukuro
	case valueobject.CampusTypeNakameguro:
		campusType = sqlc.CampusTypeNakameguro
	default:
		return nil, errs.ErrInvalidCampusType
	}

	id, err := repo.querier.CreateReservation(ctx, sqlc.CreateReservationParams{
		ID:         params.ID,
		UserID:     params.UserID,
		CampusType: campusType,
		RoomID:     params.RoomID,
		Date:       params.Date,
		FromHour:   int32(params.FromHour),
		FromMinute: int32(params.FromMinute),
		ToHour:     int32(params.ToHour),
		ToMinute:   int32(params.ToMinute),
		Note:       params.Note,
	})
	if err != nil {
		return nil, err
	}

	return &id, nil
}
