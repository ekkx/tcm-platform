package jwt

import (
	"errors"
	"time"

	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrInvalidTokenType = errors.New("invalid token type")
	ErrExpiredToken     = errors.New("expired token")
	ErrInvalidSignature = errors.New("invalid signature")
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

type Claims struct {
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secretKey         string
	accessExpiration  time.Duration
	refreshExpiration time.Duration
}

func NewJWTManager(secretKey string, accessExpiration, refreshExpiration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:         secretKey,
		accessExpiration:  accessExpiration,
		refreshExpiration: refreshExpiration,
	}
}

func (m *JWTManager) GenerateAccessToken(userID ulid.ULID) (string, error) {
	claims := &Claims{
		TokenType: TokenTypeAccess,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.accessExpiration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) GenerateRefreshToken(userID ulid.ULID) (string, error) {
	claims := &Claims{
		TokenType: TokenTypeRefresh,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.refreshExpiration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSignature
		}
		return []byte(m.secretKey), nil
	}, jwt.WithExpirationRequired())

	if err != nil {
		// なぜか errors.Is(err, jwt.ErrTokenExpired) を素直に使えないので独自で有効期限を検証する
		var claims *Claims
		if token != nil {
			if t, ok := token.Claims.(*Claims); ok {
				claims = t
			}
		}
		if claims != nil && claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	switch claims.TokenType {
	case TokenTypeAccess, TokenTypeRefresh:
		// Valid token types
	default:
		return nil, ErrInvalidTokenType
	}

	return claims, nil
}
