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

func (r *OrderItemRepository) ListByOrderID(
	ctx context.Context,
	orderID string,
) ([]*domain.OrderItem, error) {

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, order_id, name, price, quantity
		 FROM order_items
		 WHERE order_id = ?`,
		orderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.OrderItem

	for rows.Next() {
		var it domain.OrderItem
		if err := rows.Scan(
			&it.ID,
			&it.OrderID,
			&it.Name,
			&it.Price,
			&it.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, &it)
	}

	return items, nil
}
