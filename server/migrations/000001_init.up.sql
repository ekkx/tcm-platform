CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    encrypted_password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TYPE campus_type AS ENUM (
    '1', -- 中目黒
    '2' -- 池袋
);

CREATE TABLE IF NOT EXISTS reservations (
    id SERIAL PRIMARY KEY,
    external_id TEXT DEFAULT NULL UNIQUE,  -- ← ここが NULL = 確定してない
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    campus_type campus_type NOT NULL,
    room_id TEXT NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    from_hour INT NOT NULL,
    from_minute INT NOT NULL,
    to_hour INT NOT NULL,
    to_minute INT NOT NULL,
    booker_name TEXT DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_reservations_room_id_date ON reservations(room_id, date);

CREATE INDEX idx_reservations_date ON reservations(date);

CREATE INDEX idx_reservations_user_id_date ON reservations(user_id, date);
