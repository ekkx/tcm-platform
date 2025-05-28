package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/core/apperrors"
	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/output"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwter"
	"github.com/golang-jwt/jwt/v5"
)

func (uc *Usecase) Authorize(ctx context.Context, params *input.Authorize) (*output.Authorize, error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.InvalidArgument.WithCause(err)
	}

	if err := uc.tcmClient.Login(&tcmrsv.LoginParams{
		UserID:   params.UserID,
		Password: params.Password,
	}); err != nil {
		return nil, apperrors.ErrInvalidEmailOrPassword
	}

	// ユーザーが存在しない場合は新規作成
	_, err := uc.userRepo.GetUserByID(ctx, params.UserID)
	if err != nil {
		if !errors.Is(err, apperrors.ErrUserNotFound) {
			return nil, apperrors.ErrInternal.WithCause(err)
		}

		encryptedPassword, err := cryptohelper.EncryptAES(params.Password, []byte(params.PasswordAESKey))
		if err != nil {
			return nil, apperrors.ErrInternal.WithCause(err)
		}

		_, err2 := uc.userRepo.CreateUser(ctx, &user_repo.CreateUserArgs{
			ID:                params.UserID,
			EncryptedPassword: encryptedPassword,
		})
		if err2 != nil {
			return nil, apperrors.ErrInternal.WithCause(err2)
		}
	}

	accessToken, err := jwter.Generate(
		jwt.MapClaims{
			"sub":   params.UserID,
			"exp":   jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			"scope": "access",
		},
		[]byte(params.JWTSecret),
	)
	if err != nil {
		return nil, apperrors.ErrInternal.WithCause(err)
	}

	refreshToken, err := jwter.Generate(
		jwt.MapClaims{
			"sub":   params.UserID,
			"exp":   jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			"scope": "refresh",
		},
		[]byte(params.JWTSecret),
	)
	if err != nil {
		return nil, apperrors.ErrInternal.WithCause(err)
	}

	return output.NewAuthorize(
		entity.Authorization{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	), nil
}
