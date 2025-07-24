-- name: ListUnavailableRoomIDs :many
SELECT
    reservations.room_id
FROM
    reservations
WHERE
    reservations.campus_type = sqlc.arg(campus_type)::campus_type
    AND reservations.date = sqlc.arg(date)::DATE
    AND (
        (reservations.from_hour * 60 + reservations.from_minute) < (sqlc.arg(to_hour)::INT * 60 + sqlc.arg(to_minute)::INT)
        AND (reservations.to_hour * 60 + reservations.to_minute) > (sqlc.arg(from_hour)::INT * 60 + sqlc.arg(from_minute)::INT)
    );
