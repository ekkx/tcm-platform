package service

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

type Service interface {
	GetUserByID(ctx context.Context, userID ulid.ULID) (*entity.User, error)
	GetUserByOfficialSiteID(ctx context.Context, officialSiteID string) (*entity.User, error)
	ListUsersByIDs(ctx context.Context, userIDs []ulid.ULID) ([]*entity.User, error)
	ListSlaveUsers(ctx context.Context, masterUserID ulid.ULID) ([]*entity.User, error)
}

type ServiceImpl struct {
	userRepo repository.Repository
}

func New(userRepo repository.Repository) Service {
	return &ServiceImpl{
		userRepo: userRepo,
	}
}
