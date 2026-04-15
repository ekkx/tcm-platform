-- name: GetSubscriptionByUserID :one
SELECT
    *
FROM
    subscriptions
WHERE
    subscriptions.user_id = sqlc.arg(user_id)::ulid;
