-- name: InsertResult :exec
INSERT INTO results (url, status_code, duration_ms, error)
VALUES ($1, $2, $3, $4);

-- name: GetLastResults :many
SELECT url, status_code, duration_ms, error, created_at
FROM results
ORDER BY created_at DESC
    LIMIT $1;

-- name: GetResultByURL :one
SELECT url, status_code, duration_ms, error, created_at
FROM results
WHERE url = $1
ORDER BY created_at DESC
    LIMIT 1;

-- name: DeleteOldResults :exec
DELETE FROM results
WHERE created_at < NOW() - INTERVAL '30 days';

-- name: CountResults :one
SELECT COUNT(*) FROM results;
