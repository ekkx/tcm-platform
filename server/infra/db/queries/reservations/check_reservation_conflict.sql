-- name: CheckReservationConflict :one
SELECT EXISTS (
  SELECT
    1
  FROM
    reservations
  WHERE
    room_id = @room_id
    AND date = @date::timestamptz
    AND (
      ((@from_hour * 60) + @from_minute) < ((to_hour * 60) + to_minute)
      AND ((@to_hour * 60) + @to_minute) > ((from_hour * 60) + from_minute)
    )
) AS conflict;
