package domain

import "time"

type OrderStatus string

const (
	OrderStatusOpen   OrderStatus = "OPEN"
	OrderStatusClosed OrderStatus = "CLOSED"
)

type Order struct {
	ID        string
	UserID    string
	ShiftID   string
	Status    OrderStatus
	Total     int64 // cents
	CreatedAt time.Time
	ClosedAt  *time.Time
}
