package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
	userRepo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/golang-jwt/jwt/v5"
)

type AuthorizeInput struct {
	UserID         string
	Password       string
	PasswordAESKey string
	JWTSecret      string
}

type AuthorizeOutput struct {
	Authorization entity.Authorization
}

func generateToken(jwtSecret []byte, claims jwt.Claims) (string, *apperrors.Error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", apperrors.ErrInternal.WithCause(err)
	}
	return tokenString, nil
}

func (uc *Usecase) Authorize(ctx context.Context, input *AuthorizeInput) (*AuthorizeOutput, *apperrors.Error) {
	if err := uc.tcmClient.Login(&tcmrsv.LoginParams{
		UserID:   input.UserID,
		Password: input.Password,
	}); err != nil {
		return nil, apperrors.ErrUnauthorized
	}

	// ユーザーが存在しない場合は新規作成
	_, err := uc.userRepo.GetUserByID(ctx, input.UserID)
	if err != nil {
		if !errors.Is(err, apperrors.ErrUserNotFound) {
			return nil, err
		}

		encryptedPassword, err := cryptohelper.EncryptAES(input.Password, []byte(input.PasswordAESKey))
		if err != nil {
			return nil, apperrors.ErrInternal.WithCause(err)
		}

		_, err2 := uc.userRepo.CreateUser(ctx, &userRepo.CreateUserArgs{
			ID:                input.UserID,
			EncryptedPassword: encryptedPassword,
		})
		if err2 != nil {
			return nil, apperrors.ErrInternal.WithCause(err2)
		}
	}

	accessToken, err := generateToken(
		[]byte(input.JWTSecret),
		jwt.MapClaims{
			"sub":   input.UserID,
			"exp":   jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			"scope": "access",
		},
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateToken(
		[]byte(input.JWTSecret),
		jwt.MapClaims{
			"sub":   input.UserID,
			"exp":   jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			"scope": "refresh",
		},
	)
	if err != nil {
		return nil, err
	}

	return &AuthorizeOutput{
		Authorization: entity.Authorization{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
