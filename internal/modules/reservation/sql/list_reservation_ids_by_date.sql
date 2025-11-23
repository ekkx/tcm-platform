-- name: ListReservationIDsByDate :many
SELECT
    reservations.id
FROM
    reservations
WHERE
    reservations.date = sqlc.arg(date)::DATE
ORDER BY
    reservations.id DESC;
