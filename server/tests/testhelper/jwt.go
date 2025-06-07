package testhelper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateTestJWT generates a test JWT token for testing purposes
func GenerateTestJWT(userID string, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"scope": "access",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// GenerateExpiredTestJWT generates an expired test JWT token
func GenerateExpiredTestJWT(userID string, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"scope": "access",
		"exp":   time.Now().Add(-time.Hour).Unix(), // Expired 1 hour ago
		"iat":   time.Now().Add(-time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// GenerateTestJWTWithClaims generates a test JWT token with custom claims
func GenerateTestJWTWithClaims(claims jwt.MapClaims, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}