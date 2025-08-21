package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/errs"
	"github.com/ekkx/tcm-platform/internal/platform/jwt"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

func (uc *UseCaseImpl) Reauthorize(ctx context.Context, input *ReauthorizeInput) (*ReauthorizeOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	claims, err := uc.jwtManager.VerifyToken(input.RefreshToken)
	if err != nil {
		switch err {
		case jwt.ErrInvalidToken, jwt.ErrExpiredToken:
			return nil, errs.ErrInvalidToken
		default:
			return nil, errs.ErrInternal.WithCause(err)
		}
	}

	if claims.TokenType != jwt.TokenTypeRefresh {
		return nil, errs.ErrInvalidTokenType
	}

	userID, err := ulid.Parse(claims.Subject)
	if err != nil {
		// サーバー側で生成したトークンなので ID が不正なのは、サーバー内部の構造的ミス
		return nil, errs.ErrInternal.WithCause(err)
	}

	user, err := uc.userQuery.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.ErrUserNotFound
	}

	auth, err := uc.issueTokens(user.ID)
	if err != nil {
		return nil, err
	}

	return NewReauthorizeOutput(*auth), nil
}
