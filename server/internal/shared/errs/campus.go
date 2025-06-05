package errs

import "google.golang.org/grpc/codes"

var (
	ErrInvalidCampusType = &Error{
		Code:     "invalid_campus_type",
		Message:  "invalid campus type",
		GRPCCode: codes.InvalidArgument,
	}
)
