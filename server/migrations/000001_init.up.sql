CREATE DOMAIN ulid AS TEXT CHECK (LENGTH(VALUE) = 26);

CREATE TABLE IF NOT EXISTS users (
    id ulid PRIMARY KEY,
    display_name VARCHAR(255) NOT NULL DEFAULT '未設定',
    master_user_id ulid DEFAULT NULL REFERENCES users(id) ON DELETE CASCADE,
    encrypted_password TEXT NOT NULL,
    create_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TYPE campus_type AS ENUM (
    'nakameguro', -- 中目黒
    'ikebukuro' -- 池袋
);

CREATE TABLE IF NOT EXISTS reservations (
    id ulid PRIMARY KEY,
    external_id TEXT DEFAULT NULL UNIQUE,  -- ← ここが NULL = 確定してない
    user_id ulid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    campus_type campus_type NOT NULL,
    room_id TEXT NOT NULL,
    date TIMESTAMPTZ NOT NULL,
    from_hour INT NOT NULL,
    from_minute INT NOT NULL,
    to_hour INT NOT NULL,
    to_minute INT NOT NULL,
    create_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_reservations_room_id_date ON reservations(room_id, date);

CREATE INDEX idx_reservations_date ON reservations(date);

CREATE INDEX idx_reservations_user_id_date ON reservations(user_id, date);
