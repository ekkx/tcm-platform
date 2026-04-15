-- プラン定義（unlimited = テスター/友人向け無制限枠）
CREATE TYPE plan_type AS ENUM ('unlimited', 'lite', 'standard', 'pro');

-- サブスクリプションテーブル
CREATE TABLE IF NOT EXISTS subscriptions (
    id ulid PRIMARY KEY,
    user_id ulid NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    -- Stripe連携（unlimited の場合は NULL）
    stripe_customer_id TEXT UNIQUE,
    stripe_subscription_id TEXT UNIQUE,
    stripe_price_id TEXT,

    -- プラン情報
    plan plan_type NOT NULL DEFAULT 'lite',
    monthly_hours INT,  -- NULL = 無制限

    -- 状態（unlimited は常に 'active'）
    status TEXT NOT NULL DEFAULT 'active',
    current_period_start TIMESTAMPTZ,
    current_period_end TIMESTAMPTZ,

    create_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_subscriptions_user_id ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_stripe_customer_id ON subscriptions(stripe_customer_id);
