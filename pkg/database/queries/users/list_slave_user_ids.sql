-- name: ListSlaveUserIDs :many
SELECT
    users.id as user_id
FROM
    users
WHERE
    users.master_user_id = sqlc.arg(master_user_id)::ulid
ORDER BY
    users.id DESC;
