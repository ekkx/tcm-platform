package repository

import (
	"context"
	"errors"

	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
	"github.com/ekkx/tcmrsv-web/server/pkg/apperrors"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetUserByID(ctx context.Context, id string) (entity.User, *apperrors.Error) {
	row := r.db.QueryRow(ctx, `
        SELECT
            users.id,
            users.encrypted_password,
            users.created_at
        FROM
            users
        WHERE
            users.id = $1
    `, id,
	)

	var u entity.User
	err := row.Scan(&u.ID, &u.EncryptedPassword, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return u, apperrors.ErrUserNotFound
		}
		return u, apperrors.ErrInternal.WithCause(err)
	}

	return u, nil
}
