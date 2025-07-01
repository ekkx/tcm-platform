-- name: ListUsersByIDs :many
SELECT
    users.*
FROM
    users
WHERE
    users.id = ANY(sqlc.arg(user_ids)::ulid[]);
