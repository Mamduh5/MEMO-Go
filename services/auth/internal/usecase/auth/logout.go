package auth

import "context"

func (u *AuthUsecase) Logout(
	ctx context.Context,
	userID string,
) error {

	return u.tokenRepo.RevokeAllByUser(ctx, userID)
}
