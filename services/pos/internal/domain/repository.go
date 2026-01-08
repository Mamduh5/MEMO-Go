package domain

import (
	"context"
	"time"
)

type ShiftRepository interface {
	Create(ctx context.Context, shift *Shift) error
	FindOpenByUserID(ctx context.Context, userID string) (*Shift, error)
	Close(ctx context.Context, shiftID string, closedAt time.Time) error
}

type OrderRepository interface {
	Create(ctx context.Context, order *Order) error
	FindByID(ctx context.Context, orderID string) (*Order, error)
}

type OrderItemRepository interface {
	Add(ctx context.Context, item *OrderItem) error
	// ListByOrderID(ctx context.Context, orderID string) ([]*OrderItem, error)
}
