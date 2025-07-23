package errs

import "connectrpc.com/connect"

var (
	ErrInvalidCampusType = &Error{
		Code:        "invalid_campus_type",
		Message:     "invalid campus type",
		ConnectCode: connect.CodeInvalidArgument,
	}
)
