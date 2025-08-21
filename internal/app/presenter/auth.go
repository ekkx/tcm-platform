package presenter

import (
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	authv1 "github.com/ekkx/tcm-platform/internal/gen/pb/auth/v1"
)

func ToAuth(auth *entity.Auth) *authv1.Auth {
	if auth == nil {
		return nil
	}
	return &authv1.Auth{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
		UserId:       auth.UserID.String(),
	}
}
