package errs

import (
	"fmt"

	"connectrpc.com/connect"
)

type Error struct {
	Code        string
	Message     string
	Cause       error
	ConnectCode connect.Code
}

func New(code, message string, connectCode connect.Code) *Error {
	return &Error{
		Code:        code,
		Message:     message,
		ConnectCode: connectCode,
	}
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func (e *Error) Unwrap() error {
	return e.Cause
}

func (e *Error) WithMessage(message string) *Error {
	return &Error{
		Code:        e.Code,
		Message:     message,
		Cause:       e.Cause,
		ConnectCode: e.ConnectCode,
	}
}

func (e *Error) WithCause(cause error) *Error {
	return &Error{
		Code:        e.Code,
		Message:     e.Message,
		Cause:       cause,
		ConnectCode: e.ConnectCode,
	}
}
