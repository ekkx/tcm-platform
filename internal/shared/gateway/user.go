package gateway

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

type UserQuery interface {
	GetUserByID(ctx context.Context, userID ulid.ULID) (*entity.User, error)
	GetUserByOfficialSiteID(ctx context.Context, officialSiteID string) (*entity.User, error)
	ListUsersByIDs(ctx context.Context, userIDs []ulid.ULID) ([]*entity.User, error)
}

type UserCommand interface {
	CreateUser(ctx context.Context, cmd *CreateUserCommand) (*ulid.ULID, error)
	UpdateUserByID(ctx context.Context, cmd *UpdateUserCommand) error
}

type CreateUserCommand struct {
	ID                   ulid.ULID
	Password             string
	OfficialSiteID       *string
	OfficialSitePassword *string
	MasterUserID         *ulid.ULID
	DisplayName          string
}

type UpdateUserCommand struct {
	UserID               ulid.ULID
	Password             *string
	OfficialSitePassword *string
	DisplayName          *string
}
