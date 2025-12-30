package auth

import (
	"context"
	"errors"
	"time"

	"memo-go/services/auth/internal/domain"

	"github.com/google/uuid"
)

var ErrInvalidRefreshToken = errors.New("invalid refresh token")

func (u *AuthUsecase) Refresh(
	ctx context.Context,
	refreshToken string,
) (newAccess string, newRefresh string, err error) {

	// 1. Load refresh token
	rt, err := u.tokenRepo.Find(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}
	if rt == nil || rt.Revoked {
		return "", "", ErrInvalidRefreshToken
	}

	// 2. Check expiration
	if time.Now().After(rt.ExpiresAt) {
		return "", "", ErrInvalidRefreshToken
	}

	// 3. Revoke old token
	if err := u.tokenRepo.Revoke(ctx, rt.ID); err != nil {
		return "", "", err
	}

	// 4. Generate new tokens
	newAccess, err = u.tokenGen.GenerateAccessToken(rt.UserID)
	if err != nil {
		return "", "", err
	}

	newRefresh, err = u.tokenGen.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	// 5. Save new refresh token
	newRT := &domain.RefreshToken{
		ID:        uuid.NewString(),
		UserID:    rt.UserID,
		Token:     newRefresh,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		Revoked:   false,
	}

	if err := u.tokenRepo.Save(ctx, newRT); err != nil {
		return "", "", err
	}

	return newAccess, newRefresh, nil
}
