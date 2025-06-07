package errs

import "google.golang.org/grpc/codes"

var (
	ErrUnauthorized = &Error{
		Code:     "unauthorized",
		Message:  "unauthorized",
		GRPCCode: codes.Unauthenticated,
	}
	ErrInvalidAccessToken = &Error{
		Code:     "invalid_access_token",
		Message:  "invalid access token",
		GRPCCode: codes.Unauthenticated,
	}
	ErrInvalidRefreshToken = &Error{
		Code:     "invalid_refresh_token",
		Message:  "invalid refresh token",
		GRPCCode: codes.Unauthenticated,
	}
	ErrAccessTokenExpired = &Error{
		Code:     "access_token_expired",
		Message:  "access token expired",
		GRPCCode: codes.Unauthenticated,
	}
	ErrRefreshTokenExpired = &Error{
		Code:     "refresh_token_expired",
		Message:  "refresh token expired",
		GRPCCode: codes.Unauthenticated,
	}
	ErrInvalidJWTScope = &Error{
		Code:     "invalid_jwt_scope",
		Message:  "invalid jwt scope",
		GRPCCode: codes.PermissionDenied,
	}
	ErrInvalidEmailOrPassword = &Error{
		Code:     "invalid_email_or_password",
		Message:  "invalid email or password",
		GRPCCode: codes.InvalidArgument,
	}
)
