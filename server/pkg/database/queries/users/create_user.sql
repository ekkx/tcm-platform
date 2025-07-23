-- name: CreateUser :one
INSERT INTO
    users (
        id,
        password,
        official_site_id,
        official_site_password,
        master_user_id,
        display_name,
        create_time
    )
VALUES
    (
        sqlc.arg(id)::ulid,
        sqlc.arg(password)::TEXT,
        sqlc.narg(official_site_id)::TEXT,
        sqlc.narg(official_site_password)::TEXT,
        sqlc.narg(master_user_id)::ulid,
        sqlc.arg(display_name)::TEXT,
        NOW()
    )
RETURNING users.id;
