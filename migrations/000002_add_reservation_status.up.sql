CREATE TYPE reservation_status AS ENUM (
    'pending',
    'success',
    'failed'
);

ALTER TABLE reservations
    ADD COLUMN status reservation_status NOT NULL DEFAULT 'pending';
