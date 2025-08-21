package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/modules/user/repository"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

func (uc *UseCaseImpl) CreateSlaveUser(ctx context.Context, params *CreateSlaveUserInput) (*CreateSlaveUserOutput, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	if !params.Actor.IsMaster() {
		return nil, errs.ErrPermissionDenied.WithMessage("only master users can create slave users")
	}

	userID, err := uc.userRepo.CreateUser(ctx, &repository.CreateUserParams{
		Password:     params.Password,
		MasterUserID: &params.Actor.ID,
		DisplayName:  params.DisplayName,
	})
	if err != nil {
		return nil, err
	}

	v, err := uc.userAsm.Build(ctx, *userID)
	if err != nil {
		return nil, err
	}

	return NewCreateSlaveUserOutput(*v), nil
}
