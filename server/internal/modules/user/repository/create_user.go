package repository

import "context"

type CreateUserArgs struct {
	ID                string `json:"id"`
	EncryptedPassword string `json:"encrypted_password"`
}

func (r *Repository) CreateUser(ctx context.Context, args *CreateUserArgs) (string, error) {
	row := r.db.QueryRow(ctx, `
        INSERT INTO
            users (
                id,
                encrypted_password
            )
        VALUES
            (
                $1,
                $2
            )
        RETURNING
            users.id
    `, args.ID, args.EncryptedPassword)

	var userID string
	err := row.Scan(&userID)
	return userID, err
}
