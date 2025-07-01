-- name: UpdateUserByID :one
UPDATE
    users
SET
    display_name = COALESCE(sqlc.arg(display_name)::TEXT, users.display_name),
    encrypted_password = COALESCE(sqlc.arg(encrypted_password)::TEXT, users.encrypted_password),
    update_time = NOW()
WHERE
    users.id = sqlc.arg(user_id)::ulid
RETURNING 1;

