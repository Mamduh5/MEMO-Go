package auth

import (
	"memo-go/services/auth/internal/domain"
	"time"
)

type AuthUsecase struct {
	userRepo   domain.UserRepository
	tokenRepo  domain.RefreshTokenRepository
	hasher     domain.PasswordHasher
	tokenGen   domain.TokenGenerator
	refreshTTL time.Duration
}

func NewAuthUsecase(
	userRepo domain.UserRepository,
	tokenRepo domain.RefreshTokenRepository,
	hasher domain.PasswordHasher,
	tokenGen domain.TokenGenerator,
	refreshTTL time.Duration,
) *AuthUsecase {
	return &AuthUsecase{
		userRepo:   userRepo,
		tokenRepo:  tokenRepo,
		hasher:     hasher,
		tokenGen:   tokenGen,
		refreshTTL: refreshTTL,
	}
}
