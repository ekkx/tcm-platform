-- name: ListSlaveUsersMeta :many
SELECT
    users.id as user_id,
    users.master_user_id as master_user_id
FROM
    users
WHERE
    users.master_user_id = sqlc.arg(master_user_id)::ulid
    AND (
        sqlc.narg(last_user_id)::ulid IS NULL
        OR users.id < sqlc.narg(last_user_id)::ulid
    )
ORDER BY
    users.id DESC
LIMIT
    $1;
