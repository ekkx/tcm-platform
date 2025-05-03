-- name: UpdateUserPassword :one
UPDATE
    users
SET
    encrypted_password = @encrypted_password::text
WHERE
    users.id = @id::text
RETURNING 1;
