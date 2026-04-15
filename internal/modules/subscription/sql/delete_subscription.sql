-- name: DeleteSubscription :exec
DELETE FROM subscriptions WHERE id = $1::ulid;
