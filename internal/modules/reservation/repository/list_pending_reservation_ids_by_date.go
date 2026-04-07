package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
)

func (repo *RepositoryImpl) ListPendingReservationIDsByDate(ctx context.Context, date ymd.YMD) ([]ulid.ULID, error) {
	return repo.querier.ListPendingReservationIDsByDate(ctx, date)
}
