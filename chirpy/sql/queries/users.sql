-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: DeleteUserWithEmail :exec
DELETE FROM users 
WHERE email = $1;

-- name: SearchUser :one
SELECT * FROM users
WHERE email = $1;


-- name: SearchUserWithID :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users 
SET email = $2, hashed_password = $3, updated_at = $4
WHERE id = $1;
