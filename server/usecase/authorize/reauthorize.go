package authorize

import (
	"context"
	"errors"
	"time"

	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
	"github.com/ekkx/tcmrsv-web/server/pkg/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwter"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

type ReauthorizeInput struct {
	RefreshToken   string
	PasswordAESKey string
	JWTSecret      string
}

func (uc *AuthorizeUsecaseImpl) Reauthorize(ctx context.Context, input *ReauthorizeInput) (*AuthorizeOutput, error) {
	uID, err := jwter.Verify(input.RefreshToken, "refresh", []byte(input.JWTSecret))
	if err != nil {
		return nil, err
	}

	_, err = uc.querier.GetUserByID(ctx, *uID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, apperrors.ErrRequestUserNotFound
	}

	cfg := ctxhelper.GetConfig(ctx)

	accessToken, err := generateToken(
		[]byte(cfg.JWTSecret),
		jwt.MapClaims{
			"sub":   *uID,
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
			"sub":   *uID,
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
