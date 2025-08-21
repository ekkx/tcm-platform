package presenter

import (
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/internal/shared/assemble"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func fromEntityUser(e *entity.User) *userv1.User {
	if e == nil {
		return nil
	}
	return &userv1.User{
		Id:          e.ID.String(),
		DisplayName: e.DisplayName,
		MasterUser:  nil,
		CreateTime:  timestamppb.New(e.CreateTime),
	}
}

func ToUser(v *assemble.UserView) *userv1.User {
	if v == nil {
		return nil
	}

	u := fromEntityUser(&v.User)
	u.MasterUser = fromEntityUser(v.MasterUser)

	return u
}

func ToUserList(users []*assemble.UserView) []*userv1.User {
	if users == nil {
		return nil
	}
	out := make([]*userv1.User, 0, len(users))
	for _, v := range users {
		out = append(out, ToUser(v))
	}
	return out
}
