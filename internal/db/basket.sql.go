// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: basket.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const AddBasketItem = `-- name: AddBasketItem :one
INSERT INTO basket(product_id, quantity)
VALUES ($1, $2)
RETURNING id
`

type AddBasketItemParams struct {
	ProductID uuid.UUID
	Quantity  int32
}

func (q *Queries) AddBasketItem(ctx context.Context, arg AddBasketItemParams) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.addBasketItemStmt, AddBasketItem, arg.ProductID, arg.Quantity)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const DeleteBasketItem = `-- name: DeleteBasketItem :exec
DELETE FROM basket
WHERE id = $1
`

func (q *Queries) DeleteBasketItem(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteBasketItemStmt, DeleteBasketItem, id)
	return err
}
