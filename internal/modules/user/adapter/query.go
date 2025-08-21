package adapter

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/internal/shared/gateway"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
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
