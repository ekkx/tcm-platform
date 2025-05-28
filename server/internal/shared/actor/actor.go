package actor

type Role string

const (
	RoleUser   Role = "user"
	RoleSystem Role = "system"
)

type Actor struct {
	ID   string
	Role Role
}
