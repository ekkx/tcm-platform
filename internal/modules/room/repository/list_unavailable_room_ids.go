package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
)

type ListUnavailableRoomIDsParams struct {
	CampusType valueobject.CampusType
	Date       ymd.YMD
	FromHour   int
	FromMinute int
	ToHour     int
	ToMinute   int
}

func (repo *RepositoryImpl) ListUnavailableRoomIDs(ctx context.Context, params *ListUnavailableRoomIDsParams) ([]string, error) {
	var campusType sqlc.CampusType
	switch params.CampusType {
	case valueobject.CampusTypeIkebukuro:
		campusType = sqlc.CampusTypeIkebukuro
	case valueobject.CampusTypeNakameguro:
		campusType = sqlc.CampusTypeNakameguro
	default:
		return nil, errs.ErrInvalidCampusType
	}

	roomIDs, err := repo.querier.ListUnavailableRoomIDs(ctx, sqlc.ListUnavailableRoomIDsParams{
		CampusType: campusType,
		Date:       params.Date,
		FromHour:   int32(params.FromHour),
		FromMinute: int32(params.FromMinute),
		ToHour:     int32(params.ToHour),
		ToMinute:   int32(params.ToMinute),
	})
	if err != nil {
		return nil, err
	}

	return roomIDs, nil
}
