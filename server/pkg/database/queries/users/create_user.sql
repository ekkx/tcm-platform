-- name: CreateUser :one
INSERT INTO
    users (
        id,
        display_name,
        master_user_id,
        encrypted_password,
        create_time,
        update_time
    )
VALUES
    (
        sqlc.arg(id)::ulid,
        sqlc.arg(display_name)::TEXT,
        sqlc.narg(master_user_id)::ulid,
        sqlc.arg(encrypted_password)::TEXT,
        NOW(),
        NOW()
    )
RETURNING users.id;
