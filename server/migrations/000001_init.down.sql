-- drop indexes
DROP INDEX IF EXISTS idx_reservations_user_id_date;
DROP INDEX IF EXISTS idx_reservations_date;
DROP INDEX IF EXISTS idx_reservations_room_id_date;

-- drop tables
DROP TABLE IF EXISTS reservations;
DROP TABLE IF EXISTS users;

-- drop types
DROP TYPE IF EXISTS campus_type;
