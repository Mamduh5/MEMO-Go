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
	CreatedAt time.Time
}
