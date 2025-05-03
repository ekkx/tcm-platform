CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    encrypted_password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS reservations (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    room_id TEXT NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    from_hour INT NOT NULL,
    from_minute INT NOT NULL,
    to_hour INT NOT NULL,
    to_minute INT NOT NULL,
    booker_name TEXT DEFAULT NULL
);
