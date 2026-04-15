-- name: GetSubscriptionByStripeCustomerID :one
SELECT
    *
FROM
    subscriptions
WHERE
    subscriptions.stripe_customer_id = sqlc.arg(stripe_customer_id)::TEXT;
