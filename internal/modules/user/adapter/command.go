package adapter

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/internal/shared/gateway"
	"github.com/ekkx/tcmrsv-web/pkg/ulid"
)

type CommandAdapter struct {
	userRepo repository.Repository
}

func NewCommandAdapter(userRepo repository.Repository) gateway.UserCommand {
	return &CommandAdapter{
		userRepo: userRepo,
	}
}

func (c *CommandAdapter) CreateUser(ctx context.Context, cmd *gateway.CreateUserCommand) (*ulid.ULID, error) {
	return c.userRepo.CreateUser(ctx, &repository.CreateUserParams{
		ID:                   cmd.ID,
		Password:             cmd.Password,
		OfficialSiteID:       cmd.OfficialSiteID,
		OfficialSitePassword: cmd.OfficialSitePassword,
		MasterUserID:         cmd.MasterUserID,
		DisplayName:          cmd.DisplayName,
	})
}

func (c *CommandAdapter) UpdateUserByID(ctx context.Context, cmd *gateway.UpdateUserCommand) error {
	return c.userRepo.UpdateUserByID(ctx, &repository.UpdateUserByIDParams{
		UserID:               cmd.UserID,
		Password:             cmd.Password,
		OfficialSitePassword: cmd.OfficialSitePassword,
		DisplayName:          cmd.DisplayName,
	})
}
