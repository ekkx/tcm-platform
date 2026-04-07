-- name: ListPendingReservationIDsByDate :many
SELECT
    reservations.id
FROM
    reservations
WHERE
    reservations.date = sqlc.arg(date)::DATE
    AND reservations.status = 'pending'
ORDER BY
    reservations.id DESC;
