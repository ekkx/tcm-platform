package repository

import (
	"context"
	"errors"

	"github.com/ekkx/tcmrsv-web/pkg/ulid"
	"github.com/jackc/pgx/v5"
)

func (repo *RepositoryImpl) GetUserIDByOfficialSiteID(ctx context.Context, officialSiteID string) (*ulid.ULID, error) {
	id, err := repo.querier.GetUserIDByOfficialSiteID(ctx, officialSiteID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &id, nil
}
