-- name: CreateAdmin :one
INSERT INTO admin (id, name, pw, phone) VALUES ($1, $2, $3, $4)

-- name: GetAdmin :one
SELECT * FROM admin WHERE id = $1

