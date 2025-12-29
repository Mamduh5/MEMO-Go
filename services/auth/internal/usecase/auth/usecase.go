package auth

import "memo-go/services/auth/internal/domain"

type AuthUsecase struct {
	userRepo  domain.UserRepository
	tokenRepo domain.RefreshTokenRepository
	hasher    domain.PasswordHasher
	tokenGen  domain.TokenGenerator
}

func NewAuthUsecase(
	userRepo domain.UserRepository,
	tokenRepo domain.RefreshTokenRepository,
	hasher domain.PasswordHasher,
	tokenGen domain.TokenGenerator,
) *AuthUsecase {
	return &AuthUsecase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		hasher:    hasher,
		tokenGen:  tokenGen,
	}
}
