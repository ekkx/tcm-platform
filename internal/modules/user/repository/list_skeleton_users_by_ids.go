package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/internal/shared/mapper"
	"github.com/ekkx/tcmrsv-web/internal/shared/util"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

func (repo *RepositoryImpl) ListSkeletonUsersByIDs(ctx context.Context, userIDs []ulid.ULID) ([]*entity.User, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}

	dbUsers, err := repo.querier.ListUsersByIDs(ctx, util.ToULIDStrings(userIDs))
	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, 0, len(dbUsers))
	for _, u := range dbUsers {
		users = append(users, mapper.ToUser(&u))
	}

	return users, nil
}
