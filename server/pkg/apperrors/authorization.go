package apperrors

import "google.golang.org/grpc/codes"

var (
	ErrUnauthorized = &Error{
		Code:     "unauthorized",
		Message:  "unauthorized",
		GRPCCode: codes.Unauthenticated,
	}
	ErrInvalidToken = &Error{
		Code:     "invalid_token",
		Message:  "invalid token",
		GRPCCode: codes.Unauthenticated,
	}
	ErrTokenExpired = &Error{
		Code:     "token_expired",
		Message:  "token expired",
		GRPCCode: codes.Unauthenticated,
	}
	ErrInvalidTokenScope = &Error{
		Code:     "invalid_token_scope",
		Message:  "invalid token scope",
		GRPCCode: codes.PermissionDenied,
	}
	ErrInvalidEmailOrPassword = &Error{
		Code:     "invalid_email_or_password",
		Message:  "invalid email or password",
		GRPCCode: codes.Unauthenticated,
	}
)
