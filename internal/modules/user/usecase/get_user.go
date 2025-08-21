package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/internal/shared/errs"
)

func (uc *UseCaseImpl) GetUser(ctx context.Context, input *GetUserInput) (*GetUserOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	v, err := uc.userAsm.Build(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	if v == nil {
		return nil, errs.ErrUserNotFound
	}

	// ユーザー本人、またはそのユーザーのマスターユーザーのみ閲覧可能
	if v.User.ID != input.Actor.ID {
		if v.MasterUser == nil || v.MasterUser.ID != input.Actor.ID {
			return nil, errs.ErrPermissionDenied.WithMessage("you can only view your own account or that of your master user")
		}
	}

	return NewGetUserOutput(*v), nil
}
