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
	ErrInvalidAuthorizationHeader = &Error{
		Code:        "invalid_authorization_header",
		Message:     "invalid authorization header",
		ConnectCode: connect.CodeUnauthenticated,
	}
	ErrInvalidToken = &Error{
		Code:        "invalid_token",
		Message:     "invalid token",
		ConnectCode: connect.CodeUnauthenticated,
	}
	ErrInvalidTokenType = &Error{
		Code:        "invalid_token_type",
		Message:     "invalid token type",
		ConnectCode: connect.CodeUnauthenticated,
	}
	ErrExpiredToken = &Error{
		Code:        "expired_token",
		Message:     "expired token",
		ConnectCode: connect.CodeUnauthenticated,
	}
)
