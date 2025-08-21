package entity

import "github.com/ekkx/tcm-platform/internal/platform/ulid"

type Auth struct {
	AccessToken  string
	RefreshToken string
	UserID       ulid.ULID
}
