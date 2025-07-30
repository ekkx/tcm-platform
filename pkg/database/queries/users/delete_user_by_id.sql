-- name: DeleteUserByID :one
DELETE FROM
    users
WHERE
    users.id = sqlc.arg(user_id)::ulid
RETURNING 1;
