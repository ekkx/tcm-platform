-- name: UpdateUserByID :one
UPDATE
    users
SET
    password = COALESCE(sqlc.narg(password)::TEXT, users.password),
    official_site_password = COALESCE(sqlc.narg(official_site_password)::TEXT, users.official_site_password),
    display_name = COALESCE(sqlc.narg(display_name)::TEXT, users.display_name)
WHERE
    users.id = sqlc.arg(user_id)::ulid
RETURNING 1;
