-- name: GetUserByID :one
SELECT
    users.id,
    users.encrypted_password
FROM
    users
WHERE
    users.id = @id::text;
