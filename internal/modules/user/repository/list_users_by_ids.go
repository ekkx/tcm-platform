package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/modules/user/mapper"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/util"
)

func (repo *RepositoryImpl) ListUsersByIDs(ctx context.Context, userIDs []ulid.ULID) ([]*entity.User, error) {
	if len(userIDs) == 0 {
		return []*entity.User{}, nil
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
