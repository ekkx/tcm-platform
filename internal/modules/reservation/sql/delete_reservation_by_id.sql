-- name: DeleteReservationByID :one
DELETE FROM
    reservations
WHERE
    reservations.id = sqlc.arg(reservation_id)::ulid
RETURNING 1;
