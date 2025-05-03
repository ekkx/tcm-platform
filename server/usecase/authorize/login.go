package authorize

import (
	"context"
	"errors"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/infra/db"
	"github.com/ekkx/tcmrsv-web/server/pkg/cryptohelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/ctxhelper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

type LoginInput struct {
	UserID   string
	Password string
}

type LoginOutput struct {
	AccessToken  string
	RefreshToken string
}

func generateToken(jwtSecret []byte, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (uc *AuthorizeUsecaseImpl) Login(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	if err := uc.tcmClient.Login(&tcmrsv.LoginParams{
		UserID:   input.UserID,
		Password: input.Password,
	}); err != nil {
		return nil, err
	}

	_, err := uc.querier.GetUserByStudentID(ctx, input.UserID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		encryptedPassword, err := cryptohelper.EncryptAES(input.Password, []byte(ctxhelper.GetConfig(ctx).PasswordAESKey))
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

	cfg := ctxhelper.GetConfig(ctx)

	accessToken, err := generateToken(
		[]byte(cfg.JWTSecret),
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
		[]byte(cfg.JWTSecret),
		jwt.MapClaims{
			"sub":   input.UserID,
			"exp":   jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			"scope": "refresh",
		},
	)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
