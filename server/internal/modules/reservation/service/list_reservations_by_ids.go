package service

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

func (svc *ServiceImpl) ListReservationsByIDs(ctx context.Context, reservationIDs []ulid.ULID) ([]*entity.Reservation, error) {
	skeletonRsvs, err := svc.reservationRepo.ListSkeletonReservationsByIDs(ctx, reservationIDs)
	if err != nil {
		return nil, err
	}

	if len(skeletonRsvs) == 0 {
		return nil, nil
	}

	// 重複を避けるためのセットを使用
	userIDSet := make(map[ulid.ULID]struct{})
	for _, rsv := range skeletonRsvs {
		userIDSet[rsv.User.ID] = struct{}{}
	}

	userIDs := make([]ulid.ULID, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	users, err := svc.userService.ListUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	completeRsvs := make([]*entity.Reservation, 0, len(skeletonRsvs))
	for _, rsv := range skeletonRsvs {
		completeRsv := rsv
		for _, user := range users {
			if user == nil {
				continue
			}
			if user.ID == rsv.User.ID {
				completeRsv.User = *user
				break
			}
		}
		completeRsvs = append(completeRsvs, completeRsv)
	}

	return completeRsvs, nil
}
