package jwter

import (
	"errors"

	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
	"github.com/golang-jwt/jwt/v5"
)

func Verify(tokenString string, scope string, secret []byte) (uID *string, err error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return secret, nil
		},
	)

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, apperrors.ErrInvalidToken
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, apperrors.ErrTokenExpired
		default:
			return nil, apperrors.ErrInvalidToken
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if scopeClaim, ok := claims["scope"].(string); ok {
			if scopeClaim != scope {
				return nil, apperrors.ErrInvalidTokenScope
			}
		}

		if sub, ok := claims["sub"].(string); ok {
			return &sub, nil
		}
		return nil, apperrors.ErrInvalidToken
	}

	return nil, apperrors.ErrInvalidToken
}
