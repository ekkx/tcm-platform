package entity

import "time"

type User struct {
	ID                string    `json:"id"`
	EncryptedPassword string    `json:"encrypted_password"`
	CreatedAt         time.Time `json:"created_at"`
}
