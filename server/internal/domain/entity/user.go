package entity

import (
	"time"

	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

type User struct {
	ID          ulid.ULID
	DisplayName string
	MasterUser  *User
	CreateTime  time.Time
	UpdateTime  time.Time
}
