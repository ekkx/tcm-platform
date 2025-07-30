package presenter

import (
	"github.com/ekkx/tcmrsv-web/internal/domain/entity"
	userv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUser(user *entity.User) *userv1.User {
	if user == nil {
		return nil
	}
	return &userv1.User{
		Id:          user.ID.String(),
		DisplayName: user.DisplayName,
		MasterUser:  ToUser(user.MasterUser),
		CreateTime:  timestamppb.New(user.CreateTime),
	}
}

func ToUserList(users []*entity.User) []*userv1.User {
	if users == nil {
		return nil
	}
	result := make([]*userv1.User, len(users))
	for i, user := range users {
		result[i] = ToUser(user)
	}
	return result
}
