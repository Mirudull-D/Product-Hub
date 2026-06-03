-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, password
FROM users
WHERE email = $1 LIMIT 1;

-- name: GetProducts :many
SELECT *
FROM products
ORDER BY id;

-- name: GetProductsById :many
SELECT *
FROM products
WHERE id = ANY($1::int[]);

-- name: UpdateProduct :exec
UPDATE products
SET
    name = $2,
    price = $3,
    image = $4,
    description = $5,
    quantity = $6
WHERE id = $1;

-- name: CreateProducts :one
INSERT INTO products(name,description,image,quantity,price)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: CreateOrder :one
INSERT INTO orders(user_id,total,status,address)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: CreateOrderItems :one
INSERT INTO orders_items(order_id,product_id,quantity,price)
VALUES ($1,$2,$3,$4)
    RETURNING *;