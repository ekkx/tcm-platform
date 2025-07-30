package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/pkg/database"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

type Repository interface {
	GetUserIDByOfficialSiteID(ctx context.Context, officialSiteID string) (*ulid.ULID, error)
	ListSkeletonUsersByIDs(ctx context.Context, userIDs []ulid.ULID) ([]*entity.User, error)
	ListSlaveUserIDs(ctx context.Context, masterUserID ulid.ULID) ([]ulid.ULID, error)
	CreateUser(ctx context.Context, params *CreateUserParams) (*ulid.ULID, error)
	UpdateUserByID(ctx context.Context, params *UpdateUserByIDParams) error
	DeleteUserByID(ctx context.Context, userID ulid.ULID) error
}

type RepositoryImpl struct {
	querier database.Querier
}

func New(querier database.Querier) Repository {
	return &RepositoryImpl{
		querier: querier,
	}
}
