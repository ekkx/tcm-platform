package actor

import (
	"errors"

	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

var (
	ErrInvalidActorID   = errors.New("invalid actor id")
	ErrInvalidActorType = errors.New("invalid actor type")
)

type Type string

const (
	TypeUser   Type = "user"
	TypeSystem Type = "system"
)

type Role string

const (
	RoleNone   Role = "none" // system や認証前ユーザーに割り当てる初期ロール
	RoleMaster Role = "master"
	RoleSlave  Role = "slave"
)

type OfficialSiteAuth struct {
	UserID   string
	Password string
}

type Actor struct {
	ID               ulid.ULID
	Type             Type
	Role             Role
	OfficialSiteAuth *OfficialSiteAuth
}

func New(id ulid.ULID, actorType Type) *Actor {
	return &Actor{
		ID:               id,
		Type:             actorType,
		Role:             RoleNone,
		OfficialSiteAuth: nil,
	}
}

func (a *Actor) WithRole(role Role) *Actor {
	a.Role = role
	return a
}

func (a *Actor) WithOfficialSiteAuth(auth *OfficialSiteAuth) *Actor {
	a.OfficialSiteAuth = auth
	return a
}

func (a *Actor) IsUser() bool {
	return a.Type == TypeUser
}

func (a *Actor) IsSystem() bool {
	return a.Type == TypeSystem
}

func (a *Actor) IsMaster() bool {
	return a.Type == TypeUser && a.Role == RoleMaster
}

func (a *Actor) IsSlave() bool {
	return a.Type == TypeUser && a.Role == RoleSlave
}
