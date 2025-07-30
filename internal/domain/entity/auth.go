package entity

type Auth struct {
	AccessToken  string
	RefreshToken string
	User         User
}
