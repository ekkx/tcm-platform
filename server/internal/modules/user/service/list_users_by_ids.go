package service

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

func (svc *ServiceImpl) ListUsersByIDs(ctx context.Context, userIDs []ulid.ULID) ([]*entity.User, error) {
	skeletonUsers, err := svc.userRepo.ListSkeletonUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	if len(skeletonUsers) == 0 {
		return nil, nil
	}

	// 重複を避けるためのセットを使用
	masterUserIDSet := make(map[ulid.ULID]struct{})
	for _, user := range skeletonUsers {
		if user.MasterUser != nil {
			masterUserIDSet[user.MasterUser.ID] = struct{}{}
		}
	}

	masterUserIDs := make([]ulid.ULID, 0, len(masterUserIDSet))
	for id := range masterUserIDSet {
		masterUserIDs = append(masterUserIDs, id)
	}

	masterUsers, err := svc.userRepo.ListSkeletonUsersByIDs(ctx, masterUserIDs)
	if err != nil {
		return nil, err
	}

	masterUserMap := make(map[ulid.ULID]*entity.User)
	for _, masterUser := range masterUsers {
		masterUserMap[masterUser.ID] = masterUser
	}

	completeUsers := make([]*entity.User, 0, len(skeletonUsers))
	for _, user := range skeletonUsers {
		completeUser := user
		if user.MasterUser != nil {
			if masterUser, exists := masterUserMap[user.MasterUser.ID]; exists {
				completeUser.MasterUser = masterUser
			}
		}
		completeUsers = append(completeUsers, completeUser)
	}
	return completeUsers, nil
}
