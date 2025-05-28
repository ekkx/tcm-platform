package input

import (
	"context"

	"github.com/go-playground/validator/v10"

	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/authorization"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
)

type Authorize struct {
	UserID         string `validate:"required"`
	Password       string `validate:"required"`
	PasswordAESKey string `validate:"required"`
	JWTSecret      string `validate:"required"`
}

func NewAuthorize() *Authorize {
	return &Authorize{}
}

func (auth *Authorize) Validate() error {
	return validator.New().Struct(auth)
}

func (auth *Authorize) FromProto(ctx context.Context, req *auth_v1.AuthorizeRequest) *Authorize {
	auth.UserID = req.UserId
	auth.Password = req.Password
	auth.PasswordAESKey = ctxhelper.GetConfig(ctx).PasswordAESKey
	auth.JWTSecret = ctxhelper.GetConfig(ctx).JWTSecret
	return auth
}
