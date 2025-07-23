package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
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

	user, err := uc.userService.GetUserByID(ctx, *userID)
	if err != nil {
		return nil, err
	}

	return NewCreateSlaveUserOutput(*user), nil
}
