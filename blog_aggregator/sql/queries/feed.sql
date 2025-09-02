-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListFeeds :many
SELECT * FROM feeds;

-- name: FindFeed :one
SELECT * FROM feeds
WHERE url = $1;

-- name: FindFeedID :one
SELECT * FROM feeds
WHERE id = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES (
    $1, 
    $2, 
    $3, 
    $4, 
    $5
  ) RETURNING *
)
SELECT 
  iff.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feed_follow AS iff
INNER JOIN users 
  ON iff.user_id = users.id
INNER JOIN feeds 
  ON iff.feed_id = feeds.id;

-- name: GetFollowForUser :many
SELECT * FROM feed_follows 
WHERE user_id = $1;

-- name: DeleteFollow :exec
DELETE FROM feed_follows
WHERE feed_id = $1 AND user_id = $2;
