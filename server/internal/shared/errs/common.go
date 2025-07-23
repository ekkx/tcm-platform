package errs

import "connectrpc.com/connect"

var (
	ErrInvalidArgument = &Error{
		Code:        "invalid_argument",
		Message:     "invalid argument",
		ConnectCode: connect.CodeInvalidArgument,
	}
	ErrInvalidIDFormat = &Error{
		Code:        "invalid_id_format",
		Message:     "invalid id format",
		ConnectCode: connect.CodeInvalidArgument,
	}
	ErrInternal = &Error{
		Code:        "internal_error",
		Message:     "internal server error",
		ConnectCode: connect.CodeInternal,
	}
	ErrPermissionDenied = &Error{
		Code:        "permission_denied",
		Message:     "permission denied",
		ConnectCode: connect.CodePermissionDenied,
	}
)
