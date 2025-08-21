package entity

import "github.com/ekkx/tcmrsv-web/pkg/ulid"

type Auth struct {
	AccessToken  string
	RefreshToken string
	UserID       ulid.ULID
}
