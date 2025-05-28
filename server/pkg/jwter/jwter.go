package jwter

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrInvalidTokenScope = errors.New("invalid token scope")
)

func Verify(tokenString string, scope string, secret []byte) (uID string, err error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return secret, nil
		},
	)

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return "", ErrInvalidToken
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			return "", ErrTokenExpired
		default:
			return "", ErrInvalidToken
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if scopeClaim, ok := claims["scope"].(string); ok {
			if scopeClaim != scope {
				return "", ErrInvalidTokenScope
			}
		}

		if sub, ok := claims["sub"].(string); ok {
			return sub, nil
		}
		return "", ErrInvalidToken
	}

	return "", ErrInvalidToken
}

func Generate(claims jwt.Claims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
