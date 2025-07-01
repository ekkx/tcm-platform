package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/auth/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/auth/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func (uc *UsecaseImpl) Authorize(ctx context.Context, params *input.Authorize) (*output.Authorize, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	user, err := uc.userRepository.GetUserByID(ctx, params.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.ErrUserNotFound
	}

	// TODO: パスワードの検証を実装する

	return nil, nil
}
