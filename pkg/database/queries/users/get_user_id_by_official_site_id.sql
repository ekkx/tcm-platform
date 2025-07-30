-- name: GetUserIDByOfficialSiteID :one
SELECT
    users.id
FROM
    users
WHERE
    users.official_site_id = sqlc.arg(official_site_id)::TEXT;
