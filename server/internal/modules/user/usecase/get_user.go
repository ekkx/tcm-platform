package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *UseCaseImpl) GetUser(ctx context.Context, input *GetUserInput) (*GetUserOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	user, err := uc.userService.GetUserByID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.ErrUserNotFound
	}

	// ユーザー本人、またはそのユーザーのマスターユーザーのみ閲覧可能
	if user.ID != input.Actor.ID {
		if user.MasterUser == nil || user.MasterUser.ID != input.Actor.ID {
			return nil, errs.ErrPermissionDenied.WithMessage("you can only view your own account or that of your master user")
		}
	}

	return NewGetUserOutput(*user), nil
}
