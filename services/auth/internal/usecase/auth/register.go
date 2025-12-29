package auth

import (
	"context"

	"memo-go/services/auth/internal/domain"

	"github.com/google/uuid"
)

func (u *AuthUsecase) Register(
	ctx context.Context,
	email string,
	password string,
) error {

	// 1. Check if user already exists
	existing, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existing != nil {
		return ErrEmailAlreadyExists
	}

	// 2. Hash password
	hashed, err := u.hasher.Hash(password)
	if err != nil {
		return err
	}

	// 3. Create user
	user := &domain.User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: hashed,
	}

	return u.userRepo.Create(ctx, user)
}
