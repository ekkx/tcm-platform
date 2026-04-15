-- name: CreateSubscription :one
INSERT INTO
    subscriptions (
        id,
        user_id,
        stripe_customer_id,
        stripe_subscription_id,
        stripe_price_id,
        plan,
        monthly_hours,
        status,
        current_period_start,
        current_period_end,
        create_time,
        update_time
    )
VALUES
    (
        sqlc.arg(id)::ulid,
        sqlc.arg(user_id)::ulid,
        sqlc.narg(stripe_customer_id)::TEXT,
        sqlc.narg(stripe_subscription_id)::TEXT,
        sqlc.narg(stripe_price_id)::TEXT,
        sqlc.arg(plan)::plan_type,
        sqlc.narg(monthly_hours)::INT,
        sqlc.arg(status)::TEXT,
        sqlc.narg(current_period_start)::TIMESTAMPTZ,
        sqlc.narg(current_period_end)::TIMESTAMPTZ,
        NOW(),
        NOW()
    )
RETURNING subscriptions.id;
