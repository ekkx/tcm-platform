package presenter

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	authv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/auth/v1"
)

func ToAuth(auth *entity.Auth) *authv1.Auth {
	if auth == nil {
		return nil
	}
	return &authv1.Auth{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
		User:         ToUser(&auth.User),
	}
}
