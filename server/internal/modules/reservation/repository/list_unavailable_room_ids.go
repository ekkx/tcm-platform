package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/ymd"
)

type ListUnavailableRoomIDsParams struct {
	CampusType enum.CampusType
	Date       ymd.YMD
	FromHour   int
	FromMinute int
	ToHour     int
	ToMinute   int
}

func (repo *RepositoryImpl) ListUnavailableRoomIDs(ctx context.Context, params *ListUnavailableRoomIDsParams) ([]string, error) {
	var campusType database.CampusType
	switch params.CampusType {
	case enum.CampusTypeIkebukuro:
		campusType = database.CampusTypeIkebukuro
	case enum.CampusTypeNakameguro:
		campusType = database.CampusTypeNakameguro
	default:
		return nil, errs.ErrInvalidCampusType
	}

	roomIDs, err := repo.querier.ListUnavailableRoomIDs(ctx, database.ListUnavailableRoomIDsParams{
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
