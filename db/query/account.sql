-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency, created_at)
values ($1, $2, $3, $4)
RETURNING *;

-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1;

-- name: ListAccounts :many
SELECT *
FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateAccount :exec
UPDATE accounts
SET owner    = $1,
    balance  = $2,
    currency = $3
WHERE id = $4;

-- name: DeleteAccount :exec
DELETE
FROM accounts
WHERE id = $1;
