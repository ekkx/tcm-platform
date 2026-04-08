-- name: UpdateReservationNote :exec
UPDATE reservations
SET
    note = sqlc.narg(note)::TEXT
WHERE
    reservations.id = sqlc.arg(id)::ulid;
