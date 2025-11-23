package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
)

func (repo *RepositoryImpl) ListReservationIDsByDate(ctx context.Context, date ymd.YMD) ([]ulid.ULID, error) {
	if date.IsZero() || !date.IsValid() {
		return nil, nil
	}

	ids, err := repo.querier.ListReservationIDsByDate(ctx, date)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, nil
	}

	return ids, nil
}
