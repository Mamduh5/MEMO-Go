package auth

import (
	"context"
	"errors"
	"time"

	"memo-go/services/auth/internal/domain"

	"github.com/google/uuid"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

func (u *AuthUsecase) Login(
	ctx context.Context,
	email string,
	password string,
) (accessToken string, refreshToken string, err error) {

	// 1. Find user
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", ErrInvalidCredentials
	}

	// 2. Compare password
	if err := u.hasher.Compare(user.Password, password); err != nil {
		return "", "", ErrInvalidCredentials
	}

	// 3. Generate access token
	accessToken, err = u.tokenGen.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}

	// 4. Generate refresh token
	refreshToken, err = u.tokenGen.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	// 5. Persist refresh token
	rt := &domain.RefreshToken{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		Revoked:   false,
	}

	if err := u.tokenRepo.Save(ctx, rt); err != nil {
		return "", "", err
	}

	// 6. Return tokens
	return accessToken, refreshToken, nil
}
