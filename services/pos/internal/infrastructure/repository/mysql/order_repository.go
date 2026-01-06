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
