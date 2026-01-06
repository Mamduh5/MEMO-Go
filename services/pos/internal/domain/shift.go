package domain

import "time"

type Shift struct {
	ID       string
	UserID   string
	OpenedAt time.Time
	ClosedAt *time.Time
}
