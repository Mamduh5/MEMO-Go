package domain

import "context"

// repository.go
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
}

type RefreshTokenRepository interface {
	Save(ctx context.Context, token *RefreshToken) error
	Find(ctx context.Context, token string) (*RefreshToken, error)
	Revoke(ctx context.Context, tokenID string) error
	RevokeAllByUser(ctx context.Context, userID string) error
}

type TokenGenerator interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken() (string, error)
}
