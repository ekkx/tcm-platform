package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type Repository interface {
	GetUserByID(ctx context.Context, userID ulid.ULID) (*entity.User, error)
	GetUserIDByOfficialSiteID(ctx context.Context, officialSiteID string) (*ulid.ULID, error)
	ListUsersByIDs(ctx context.Context, userIDs []ulid.ULID) ([]*entity.User, error)
	ListSlaveUserIDs(ctx context.Context, masterUserID ulid.ULID) ([]ulid.ULID, error)
	CreateUser(ctx context.Context, params *CreateUserParams) (*ulid.ULID, error)
	UpdateUserByID(ctx context.Context, params *UpdateUserByIDParams) error
	DeleteUserByID(ctx context.Context, userID ulid.ULID) error
}

type RepositoryImpl struct {
	querier sqlc.Querier
}

func New(querier sqlc.Querier) Repository {
	return &RepositoryImpl{
		querier: querier,
	}
}
