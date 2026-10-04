-- name: CreateTodo :one
INSERT INTO todos (title) VALUES ($1) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1;

-- name: ListTodos :many
SELECT * FROM todos ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2, is_completed = $3, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :one
UPDATE todos
SET is_deleted = true, updated_at = now()
WHERE id = $1
RETURNING *;
