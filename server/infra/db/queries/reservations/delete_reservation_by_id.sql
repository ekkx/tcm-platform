-- name: DeleteReservationByID :one
DELETE FROM
    reservations
WHERE
    id = @reservation_id::int
RETURNING 1;
