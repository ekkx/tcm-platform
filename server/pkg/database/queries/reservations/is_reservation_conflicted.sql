-- name: IsReservationConflicted :one
SELECT EXISTS (
    SELECT
        1
    FROM
        reservations
    WHERE
        reservations.room_id = sqlc.arg(room_id)::TEXT
        AND reservations.date = sqlc.arg(date)::DATE
        AND (
            ((sqlc.arg(from_hour) * 60) + sqlc.arg(from_minute)) < ((reservations.to_hour * 60) + reservations.to_minute)
            AND ((sqlc.arg(to_hour) * 60) + sqlc.arg(to_minute)) > ((reservations.from_hour * 60) + reservations.from_minute)
        )
) AS conflict;
