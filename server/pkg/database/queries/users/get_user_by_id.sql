-- name: GetUserByID :one
SELECT
    *
FROM
    users
WHERE
    users.id = sqlc.arg(id)::ulid;
