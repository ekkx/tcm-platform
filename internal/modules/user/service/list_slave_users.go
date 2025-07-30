package service

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

func (svc *ServiceImpl) ListSlaveUsers(ctx context.Context, masterUserID ulid.ULID) ([]*entity.User, error) {
	slaveIDs, err := svc.userRepo.ListSlaveUserIDs(ctx, masterUserID)
	if err != nil {
		return nil, err
	}

	if len(slaveIDs) == 0 {
		return nil, nil
	}

	return svc.ListUsersByIDs(ctx, slaveIDs)
}
