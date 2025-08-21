package adapter

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/app/gateway"
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/modules/user/repository"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type QueryAdapter struct {
	userRepo repository.Repository
}

func NewQueryAdapter(userRepo repository.Repository) gateway.UserQuery {
	return &QueryAdapter{
		userRepo: userRepo,
	}
}

func (q *QueryAdapter) GetUserByID(ctx context.Context, userID ulid.ULID) (*entity.User, error) {
	return q.userRepo.GetUserByID(ctx, userID)
}

func (q *QueryAdapter) GetUserByOfficialSiteID(ctx context.Context, officialSiteID string) (*entity.User, error) {
	id, err := q.userRepo.GetUserIDByOfficialSiteID(ctx, officialSiteID)
	if err != nil {
		return nil, err
	}
	if id == nil {
		return nil, nil
	}
	return q.userRepo.GetUserByID(ctx, *id)
}

func (q *QueryAdapter) ListUsersByIDs(ctx context.Context, userIDs []ulid.ULID) ([]*entity.User, error) {
	return q.userRepo.ListUsersByIDs(ctx, userIDs)
}
