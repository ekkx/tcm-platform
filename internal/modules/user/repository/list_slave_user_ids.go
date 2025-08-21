package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

func (repo *RepositoryImpl) ListSlaveUserIDs(ctx context.Context, masterUserID ulid.ULID) ([]ulid.ULID, error) {
	ids, err := repo.querier.ListSlaveUserIDs(ctx, masterUserID)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
