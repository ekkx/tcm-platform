package entity

import (
	"time"

	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

type User struct {
	ID                   ulid.ULID
	Password             string
	OfficialSiteID       *string
	OfficialSitePassword *string
	MasterUser           *User
	DisplayName          string
	CreateTime           time.Time
}

func (u *User) IsMaster() bool {
	return u.MasterUser == nil
}

func (u *User) IsSlave() bool {
	return u.MasterUser != nil
}

func (u *User) CheckPassword(password string) bool {
	// TODO: パスワードのハッシュ化と比較を行う
	return u.Password == password
}
