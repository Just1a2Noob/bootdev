-- name: MarkFeedFetched :one
UPDATE feeds
SET updated_at = $1,
    last_fetched_at = $2
WHERE ID = $3
RETURNING *;

-- name: GetNextFeedtoFetch :one
SELECT * FROM feeds
ORDER BY updated_at NULLS FIRST
LIMIT 1;
