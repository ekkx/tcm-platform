package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/shared/errs"
)

func (uc *UseCaseImpl) ListSlaveUsers(ctx context.Context, input *ListSlaveUsersInput) (*ListSlaveUsersOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	if !input.Actor.IsMaster() {
		return nil, errs.ErrPermissionDenied.WithMessage("only master users can list slave users")
	}

	slaveUserIDs, err := uc.userRepo.ListSlaveUserIDs(ctx, input.Actor.ID)
	if err != nil {
		return nil, err
	}

	v, err := uc.userAsm.BuildList(ctx, slaveUserIDs)
	if err != nil {
		return nil, err
	}

	return NewListSlaveUsersOutput(v), nil
}
