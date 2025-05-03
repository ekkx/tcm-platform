package apperrors

import (
	"github.com/ekkx/tcmrsv-web/server/pkg/response"
	"github.com/labstack/echo/v4"
)

type ErrorMessage struct {
	Message any `json:"message"`
}

func ErrorHandler(err error, ctx echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		response.JSON(ctx, he.Code, &ErrorMessage{
			Message: he.Message,
		})
		return
	}

	code := getErrorCode(err)

	response.JSON(ctx, code, &ErrorMessage{
		Message: err.Error(),
	})
}
