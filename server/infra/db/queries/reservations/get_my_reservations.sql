-- name: GetMyReservations :many
SELECT
    reservations.*
FROM
    reservations
WHERE
    reservations.user_id = @user_id::text
    AND reservations.date >= @date::timestamptz;
