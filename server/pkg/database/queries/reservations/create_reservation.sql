-- name: CreateReservation :one
INSERT INTO
    reservations (
        id,
        user_id,
        campus_type,
        room_id,
        date,
        from_hour,
        from_minute,
        to_hour,
        to_minute,
        create_time
    )
VALUES
    (
        sqlc.arg(id)::ulid,
        sqlc.arg(user_id)::ulid,
        sqlc.arg(campus_type)::campus_type,
        sqlc.arg(room_id)::TEXT,
        sqlc.arg(date)::DATE,
        sqlc.arg(from_hour)::INT,
        sqlc.arg(from_minute)::INT,
        sqlc.arg(to_hour)::INT,
        sqlc.arg(to_minute)::INT,
        NOW()
    )
RETURNING reservations.id;
