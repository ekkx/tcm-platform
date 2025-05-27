package input

import (
	"context"

	auth_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/authorization"
	"github.com/ekkx/tcmrsv-web/server/pkg/ctxhelper"
	"github.com/go-playground/validator/v10"
)

type Reauthorize struct {
	RefreshToken   string `validate:"required"`
	PasswordAESKey string `validate:"required"`
	JWTSecret      string `validate:"required"`
}

func NewReauthorize() *Reauthorize {
	return &Reauthorize{}
}

func (reauth *Reauthorize) Validate() error {
	return validator.New().Struct(reauth)
}

func (reauth *Reauthorize) FromProto(ctx context.Context, req *auth_v1.ReauthorizeRequest) *Reauthorize {
	reauth.RefreshToken = req.RefreshToken
	reauth.PasswordAESKey = ctxhelper.GetConfig(ctx).PasswordAESKey
	reauth.JWTSecret = ctxhelper.GetConfig(ctx).JWTSecret
	return reauth
}
