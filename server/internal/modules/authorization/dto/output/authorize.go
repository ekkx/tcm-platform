package output

import (
	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/authorization"
	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
)

type Authorize struct {
	Authorization entity.Authorization
}

func NewAuthorize(authorization entity.Authorization) *Authorize {
	return &Authorize{
		Authorization: authorization,
	}
}

func (auth *Authorize) ToProto() *auth_v1.AuthorizeReply {
	return &auth_v1.AuthorizeReply{
		Authorization: &auth_v1.Authorization{
			AccessToken:  auth.Authorization.AccessToken,
			RefreshToken: auth.Authorization.RefreshToken,
		},
	}
}
