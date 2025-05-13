package authorize

import (
	"context"
	"errors"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/infra/db"
	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

type AuthorizeInput struct {
	UserID         string
	Password       string
	PasswordAESKey string
	JWTSecret      string
}

type AuthorizeOutput struct {
	AccessToken  string
	RefreshToken string
}

func generateToken(jwtSecret []byte, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (uc *AuthorizeUsecaseImpl) Authorize(ctx context.Context, input *AuthorizeInput) (*AuthorizeOutput, error) {
	if err := uc.tcmClient.Login(&tcmrsv.LoginParams{
		UserID:   input.UserID,
		Password: input.Password,
	}); err != nil {
		return nil, apperrors.ErrUnauthorized
	}

    // ユーザーが存在しない場合は新規作成
	_, err := uc.querier.GetUserByID(ctx, input.UserID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		encryptedPassword, err := cryptohelper.EncryptAES(input.Password, []byte(input.PasswordAESKey))
		if err != nil {
			return nil, err
		}

		_, err2 := uc.querier.CreateUser(ctx, db.CreateUserParams{
			ID:                input.UserID,
			EncryptedPassword: encryptedPassword,
		})
		if err2 != nil {
			return nil, err2
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
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
