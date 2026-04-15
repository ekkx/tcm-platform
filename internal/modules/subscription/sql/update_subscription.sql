-- name: UpdateSubscription :one
UPDATE
    subscriptions
SET
    stripe_subscription_id = COALESCE(sqlc.narg(stripe_subscription_id)::TEXT, subscriptions.stripe_subscription_id),
    stripe_price_id = COALESCE(sqlc.narg(stripe_price_id)::TEXT, subscriptions.stripe_price_id),
    plan = COALESCE(sqlc.narg(plan)::plan_type, subscriptions.plan),
    monthly_hours = COALESCE(sqlc.narg(monthly_hours)::INT, subscriptions.monthly_hours),
    status = COALESCE(sqlc.narg(status)::TEXT, subscriptions.status),
    current_period_start = COALESCE(sqlc.narg(current_period_start)::TIMESTAMPTZ, subscriptions.current_period_start),
    current_period_end = COALESCE(sqlc.narg(current_period_end)::TIMESTAMPTZ, subscriptions.current_period_end),
    update_time = NOW()
WHERE
    subscriptions.id = sqlc.arg(id)::ulid
RETURNING 1;
