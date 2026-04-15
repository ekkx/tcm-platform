-- name: GetUsedMinutesByUserID :one
SELECT
    COALESCE(SUM(
        (reservations.to_hour * 60 + reservations.to_minute) - (reservations.from_hour * 60 + reservations.from_minute)
    ), 0)::INT as used_minutes
FROM
    reservations
JOIN
    subscriptions ON subscriptions.user_id = reservations.user_id
WHERE
    reservations.user_id = sqlc.arg(user_id)::ulid
    AND reservations.status IN ('pending', 'success')
    AND reservations.create_time >= COALESCE(subscriptions.current_period_start, date_trunc('month', CURRENT_DATE))
    AND reservations.create_time < COALESCE(subscriptions.current_period_end, date_trunc('month', CURRENT_DATE) + interval '1 month');
