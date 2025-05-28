package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/apperrors"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwter"
	"github.com/golang-jwt/jwt/v5"
)

func (uc *Usecase) Reauthorize(ctx context.Context, params *input.Reauthorize) (*output.Reauthorize, error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.InvalidArgument.WithCause(err)
	}

	// リフレッシュトークンの検証
	uID, err := jwter.Verify(params.RefreshToken, "refresh", []byte(params.JWTSecret))
	if err != nil {
		switch {
		case errors.Is(err, jwter.ErrInvalidToken):
			return nil, apperrors.ErrInvalidRefreshToken
		case errors.Is(err, jwter.ErrTokenExpired):
			return nil, apperrors.ErrRefreshTokenExpired
		case errors.Is(err, jwter.ErrInvalidTokenScope):
			return nil, apperrors.ErrInvalidJWTScope
		default:
			return nil, apperrors.ErrInternal.WithCause(err)
		}
	}

	// ユーザーが存在するか確認
	u, err := uc.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			return nil, apperrors.ErrRequestUserNotFound
		}
		return nil, apperrors.ErrInternal.WithCause(err)
	}

	// 念の為TCMにログイン
	rawPassword, err := cryptohelper.DecryptAES(u.EncryptedPassword, []byte(params.PasswordAESKey))
	if err != nil {
		return nil, apperrors.ErrInternal.WithCause(err)
	}

	if err := uc.tcmClient.Login(&tcmrsv.LoginParams{
		UserID:   u.ID,
		Password: rawPassword,
	}); err != nil {
		return nil, apperrors.ErrInvalidEmailOrPassword
	}

	// アクセストークンとリフレッシュトークンを生成
	accessToken, err := jwter.Generate(
		jwt.MapClaims{
			"sub":   uID,
			"scope": "access",
			"exp":   jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
		[]byte(params.JWTSecret),
	)
	if err != nil {
		return nil, apperrors.ErrInternal.WithCause(err)
	}

	refreshToken, err := jwter.Generate(
		jwt.MapClaims{
			"sub":   uID,
			"exp":   jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			"scope": "refresh",
		},
		[]byte(params.JWTSecret),
	)
	if err != nil {
		return nil, apperrors.ErrInternal.WithCause(err)
	}

	return output.NewReauthorize(
		entity.Authorization{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	), nil
}
