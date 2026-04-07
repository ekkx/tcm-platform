-- name: UpdateReservationStatus :exec
UPDATE reservations
SET
    status = sqlc.arg(status)::reservation_status,
    official_site_id = sqlc.narg(official_site_id)::TEXT
WHERE
    reservations.id = sqlc.arg(id)::ulid;
