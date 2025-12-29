package auth

import "context"

func (u *AuthUsecase) Refresh(
	ctx context.Context,
	refreshToken string,
) (newAccess string, newRefresh string, err error) {

	// implementation later
	return
}
