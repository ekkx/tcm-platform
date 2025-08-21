-- name: ListUserReservationIDs :many
SELECT
    reservations.id
FROM
    reservations
WHERE
    reservations.user_id = sqlc.arg(user_id)::ulid
    AND reservations.date >= sqlc.arg(date)::DATE
ORDER BY
    reservations.id DESC;
