-- name: CreateAdmin :exec
INSERT INTO admin (id, name, pw, phone) VALUES (?, ?, ?, ?);

-- name: GetAdmin :one
SELECT * FROM admin WHERE id = ? LIMIT 1;

-- name: UpdateAdminPw :exec
UPDATE admin SET pw = ? WHERE id = ?;

-- name: CreateProduct :exec
INSERT INTO product (category, price, cost, name, name_initial, description, barcode, expire, size) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetProduct :one
SELECT * FROM product WHERE seq = ? LIMIT 1;

-- name: ListProductsByName :many
SELECT * FROM product WHERE name LIKE ? LIMIT ?;

-- name: ListProductsByNameInitial :many
SELECT * FROM product WHERE name_initial LIKE ? LIMIT ?;

-- name: UpdateProductIfNotNil :exec
UPDATE product
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

-- name: DeleteProduct :exec
DELETE FROM product WHERE seq = ?;