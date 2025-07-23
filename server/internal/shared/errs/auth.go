package errs

import (
	"connectrpc.com/connect"
)

var (
	ErrUnauthorized = &Error{
		Code:        "unauthorized",
		Message:     "unauthorized",
		ConnectCode: connect.CodeUnauthenticated,
	}
	ErrInvalidAccessToken = &Error{
		Code:        "invalid_access_token",
		Message:     "invalid access token",
		ConnectCode: connect.CodeUnauthenticated,
	}
	ErrInvalidRefreshToken = &Error{
		Code:        "invalid_refresh_token",
		Message:     "invalid refresh token",
		ConnectCode: connect.CodeUnauthenticated,
	}
	ErrAccessTokenExpired = &Error{
		Code:        "access_token_expired",
		Message:     "access token expired",
		ConnectCode: connect.CodeUnauthenticated,
	}
	ErrRefreshTokenExpired = &Error{
		Code:        "refresh_token_expired",
		Message:     "refresh token expired",
		ConnectCode: connect.CodeUnauthenticated,
	}
	ErrInvalidJWTScope = &Error{
		Code:        "invalid_jwt_scope",
		Message:     "invalid jwt scope",
		ConnectCode: connect.CodePermissionDenied,
	}
	ErrInvalidEmailOrPassword = &Error{
		Code:        "invalid_email_or_password",
		Message:     "invalid email or password",
		ConnectCode: connect.CodeInvalidArgument,
	}
)
