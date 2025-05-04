-- name: CreateReservation :one
INSERT INTO
    reservations (
        user_id,
        campus,
        room_id,
        date,
        from_hour,
        from_minute,
        to_hour,
        to_minute,
        booker_name
    )
VALUES
    (
        @user_id::text,
        @campus::campus,
        @room_id::text,
        @date::timestamptz,
        @from_hour::int,
        @from_minute::int,
        @to_hour::int,
        @to_minute::int,
        sqlc.narg(booker_name)::text
    )
RETURNING
    reservations.*;
