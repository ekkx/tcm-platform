-- name: CreateUser :one
INSERT INTO
    users (
        id,
        encrypted_password
    )
VALUES
    (
        @id::text,
        @encrypted_password::text
    )
RETURNING
    users.id;
