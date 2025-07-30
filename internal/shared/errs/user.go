package errs

import "connectrpc.com/connect"

var (
	ErrUserNotFound = &Error{
		Code:        "user_not_found",
		Message:     "user not found",
		ConnectCode: connect.CodeNotFound,
	}
	ErrRequestUserNotFound = &Error{
		Code:        "request_user_not_found",
		Message:     "request user not found",
		ConnectCode: connect.CodeNotFound,
	}
)
