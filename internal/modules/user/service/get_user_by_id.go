package service

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

func (svc *ServiceImpl) GetUserByID(ctx context.Context, userID ulid.ULID) (*entity.User, error) {
	users, err := svc.ListUsersByIDs(ctx, []ulid.ULID{userID})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return users[0], nil
}
