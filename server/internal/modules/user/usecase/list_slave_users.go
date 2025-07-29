package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *UseCaseImpl) ListSlaveUsers(ctx context.Context, input *ListSlaveUsersInput) (*ListSlaveUsersOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	if !input.Actor.IsMaster() {
		return nil, errs.ErrPermissionDenied.WithMessage("only master users can list slave users")
	}

	users, err := uc.userService.ListSlaveUsers(ctx, input.Actor.ID)
	if err != nil {
		return nil, err
	}

	return NewListSlaveUsersOutput(users), nil
}
