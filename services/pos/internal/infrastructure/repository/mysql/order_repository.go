package mysql

import (
	"context"
	"database/sql"

	"memo-go/services/pos/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(ctx context.Context, o *domain.Order) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO orders (id, user_id, shift_id, status, created_at)
		 VALUES (?, ?, ?, ?, ?)`,
		o.ID, o.UserID, o.ShiftID, o.Status, o.CreatedAt,
	)
	return err
}

func (r *OrderRepository) FindByID(
	ctx context.Context,
	orderID string,
) (*domain.Order, error) {

	row := r.db.QueryRowContext(
		ctx,
		`SELECT id, user_id, shift_id, status, created_at
		 FROM orders
		 WHERE id = ?`,
		orderID,
	)

	var o domain.Order

	err := row.Scan(
		&o.ID,
		&o.UserID,
		&o.ShiftID,
		&o.Status,
		&o.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &o, nil
}
