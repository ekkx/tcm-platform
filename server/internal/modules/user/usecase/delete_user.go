package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *UseCaseImpl) DeleteUser(ctx context.Context, input *DeleteUserInput) (*DeleteUserOutput, error) {
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

	// ユーザー本人、またはそのユーザーのマスターユーザーのみ削除可能
	if user.ID != input.Actor.ID {
		if user.MasterUser == nil || user.MasterUser.ID != input.Actor.ID {
			return nil, errs.ErrPermissionDenied.WithMessage("you can only delete your own account or that of your master user")
		}
	}

	if err := uc.userRepo.DeleteUserByID(ctx, input.UserID); err != nil {
		return nil, err
	}

	return NewDeleteUserOutput(), nil
}
