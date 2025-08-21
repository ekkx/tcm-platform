-- name: ListReservationsByIDs :many
SELECT
    reservations.*
FROM
    reservations
WHERE
    reservations.id = ANY(sqlc.arg(reservation_ids)::TEXT[]);
