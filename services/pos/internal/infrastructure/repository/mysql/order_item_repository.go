package mysql

import (
	"context"
	"database/sql"
	"memo-go/services/pos/internal/domain"
)

type OrderItemRepository struct {
	db *sql.DB
}

func NewOrderItemRepository(db *sql.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r *OrderItemRepository) Add(ctx context.Context, i *domain.OrderItem) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO order_items (id, order_id, name, price, quantity)
		 VALUES (?, ?, ?, ?, ?)`,
		i.ID, i.OrderID, i.Name, i.Price, i.Quantity,
	)
	return err
}
