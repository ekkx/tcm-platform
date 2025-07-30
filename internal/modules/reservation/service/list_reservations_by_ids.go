package service

import (
	"context"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/internal/shared/mapper"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
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

	rooms := tcmrsv.New().GetRooms()

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
		for _, room := range rooms {
			if room.ID == rsv.Room.ID {
				completeRsv.Room = *mapper.ToRoom(&room)
				break
			}
		}
		completeRsvs = append(completeRsvs, completeRsv)
	}

	return completeRsvs, nil
}
