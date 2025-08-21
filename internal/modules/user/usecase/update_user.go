package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/modules/user/repository"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

func (uc *UseCaseImpl) UpdateUser(ctx context.Context, input *UpdateUserInput) (*UpdateUserOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	if err := uc.userRepo.UpdateUserByID(ctx, &repository.UpdateUserByIDParams{
		UserID:      input.Actor.ID,
		DisplayName: &input.DisplayName,
	}); err != nil {
		return nil, err
	}

	user, err := uc.userAsm.Build(ctx, input.Actor.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.ErrUserNotFound
	}

	return NewUpdateUserOutput(*user), nil
}
