-- name: CreateAdmin :exec
INSERT INTO admin (id, name, pw, phone) VALUES (?, ?, ?, ?);

-- name: GetAdmin :one
SELECT * FROM admin WHERE id = ? LIMIT 1;

-- name: UpdateAdminPw :exec
UPDATE admin SET pw = ? WHERE id = ?;

-- name: CreateMenu :exec
INSERT INTO menu (category, price, cost, name, name_initial, description, barcode, expire, size) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetMenu :one
SELECT * FROM menu WHERE seq = ? LIMIT 1;

-- name: ListMenus :many
SELECT * FROM menu WHERE seq >= ? LIMIT 10;

-- name: SearchMenusByName :many
SELECT * FROM menu WHERE name LIKE ? LIMIT ?;

-- name: SearchMenusByNameInitial :many
SELECT * FROM menu WHERE name_initial LIKE ? LIMIT ?;

-- name: UpdateMenuIfNotNil :exec
UPDATE menu
SET category = COALESCE(sqlc.narg('category'), category),
    price = COALESCE(sqlc.narg('price'), price),
    cost = COALESCE(sqlc.narg('cost'), cost),
    name = COALESCE(sqlc.narg('name'), name),
    name_initial = COALESCE(sqlc.narg('name_initial'), name_initial),
    description = COALESCE(sqlc.narg('decription'), description),
    barcode = COALESCE(sqlc.narg('barcode'), barcode),
    expire = COALESCE(sqlc.narg('expire'), expire),
    size = COALESCE(sqlc.narg('size'), size)
WHERE seq = ?;

-- name: DeleteMenu :exec
DELETE FROM menu WHERE seq = ?;