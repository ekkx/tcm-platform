-- name: GetUserMetaByID :one
SELECT
    users.id AS user_id,
    users.master_user_id AS master_user_id
FROM
    users
WHERE
    users.id = sqlc.arg(user_id)::ulid;
