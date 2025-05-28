package output

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/authorization"
)

type Reauthorize struct {
	Authorization entity.Authorization
}

func NewReauthorize(authorization entity.Authorization) *Reauthorize {
	return &Reauthorize{
		Authorization: authorization,
	}
}

func (auth *Reauthorize) ToProto() *auth_v1.ReauthorizeReply {
	return &auth_v1.ReauthorizeReply{
		Authorization: &auth_v1.Authorization{
			AccessToken:  auth.Authorization.AccessToken,
			RefreshToken: auth.Authorization.RefreshToken,
		},
	}
}
