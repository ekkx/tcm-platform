package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
)

type ListUserReservationIDsParams struct {
	UserID ulid.ULID
	Date   ymd.YMD
}

func (repo *RepositoryImpl) ListUserReservationIDs(ctx context.Context, params *ListUserReservationIDsParams) ([]ulid.ULID, error) {
	if params == nil {
		return nil, nil
	}

	ids, err := repo.querier.ListUserReservationIDs(ctx, sqlc.ListUserReservationIDsParams{
		UserID: params.UserID,
		Date:   params.Date,
	})
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, nil
	}

	return ids, nil
}
