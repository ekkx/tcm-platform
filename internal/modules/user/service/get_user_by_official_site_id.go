package service

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

func (svc *ServiceImpl) GetUserByOfficialSiteID(ctx context.Context, officialSiteID string) (*entity.User, error) {
	userID, err := svc.userRepo.GetUserIDByOfficialSiteID(ctx, officialSiteID)
	if err != nil {
		return nil, err
	}

	if userID == nil {
		return nil, nil
	}

	users, err := svc.ListUsersByIDs(ctx, []ulid.ULID{*userID})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users[0], nil
}
