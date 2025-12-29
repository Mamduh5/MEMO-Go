package domain

import "time"

// user.go
type User struct {
	ID       string
	Email    string
	Password string // hashed
}

type RefreshToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	Revoked   bool
}
