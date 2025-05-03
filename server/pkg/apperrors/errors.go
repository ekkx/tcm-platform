package apperrors

import "errors"

var (
	// auth
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrInvalidTokenScope = errors.New("invalid token scope")

	// user
	ErrUserNotFound        = errors.New("user not found")
	ErrRequestUserNotFound = errors.New("request user not found")
)
