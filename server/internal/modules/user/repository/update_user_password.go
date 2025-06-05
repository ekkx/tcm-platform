package repository

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

type UpdateUserPasswordArgs struct {
	EncryptedPassword string
	ID                string
}

func (r *Repository) UpdateUserPassword(ctx context.Context, args *UpdateUserPasswordArgs) error {
	row := r.db.QueryRow(ctx, `
        UPDATE
            users
        SET
            encrypted_password = $1
        WHERE
            users.id = $2
        RETURNING 1
    `, args.EncryptedPassword, args.ID,
	)

	var n int
	err := row.Scan(&n)
	if err != nil {
		return errs.ErrInternal.WithCause(err)
	}

	return nil
}
