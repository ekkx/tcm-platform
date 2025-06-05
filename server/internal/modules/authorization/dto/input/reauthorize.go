package input

import (
	"context"

	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/authorization"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
)

type Reauthorize struct {
	RefreshToken   string `validate:"required,jwt"`
	PasswordAESKey string `validate:"required"`
	JWTSecret      string `validate:"required"`
}

func NewReauthorize() *Reauthorize {
	return &Reauthorize{}
}

func (reauth *Reauthorize) Validate() error {
	return validate.Struct(reauth)
}

func (reauth *Reauthorize) FromProto(ctx context.Context, req *auth_v1.ReauthorizeRequest) *Reauthorize {
	reauth.RefreshToken = req.RefreshToken
	reauth.PasswordAESKey = ctxhelper.GetConfig(ctx).PasswordAESKey
	reauth.JWTSecret = ctxhelper.GetConfig(ctx).JWTSecret
	return reauth
}
