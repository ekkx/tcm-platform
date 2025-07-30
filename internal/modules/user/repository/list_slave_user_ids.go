package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

func (repo *RepositoryImpl) ListSlaveUserIDs(ctx context.Context, masterUserID ulid.ULID) ([]ulid.ULID, error) {
	ids, err := repo.querier.ListSlaveUserIDs(ctx, masterUserID)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, nil
	}

	return ids, nil
}
