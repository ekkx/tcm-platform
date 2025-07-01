package output

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	authv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/auth/v1"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Authorize struct {
	entity.Auth
	entity.User
}

func NewAuthorize(auth entity.Auth, user entity.User) *Authorize {
	return &Authorize{
		Auth: auth,
		User: user,
	}
}

func (st *Authorize) ToResponse() *connect.Response[authv1.AuthorizeResponse] {
	return connect.NewResponse(&authv1.AuthorizeResponse{
		Auth: &authv1.Auth{
			AccessToken:  st.AccessToken,
			RefreshToken: st.RefreshToken,
		},
		User: &userv1.User{ // TODO: refactor to convert in presentation layer
			Id:          st.User.ID.String(),
			DisplayName: st.User.DisplayName,
			MasterUser: &userv1.User{
				Id:          st.User.MasterUser.ID.String(),
				DisplayName: st.User.MasterUser.DisplayName,
				MasterUser:  nil,
				CreateTime:  timestamppb.New(st.User.MasterUser.CreateTime),
			},
			CreateTime: timestamppb.New(st.User.CreateTime),
		},
	})
}
