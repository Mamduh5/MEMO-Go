package auth

import "context"

func (u *AuthUsecase) Login(
	ctx context.Context,
	email string,
	password string,
) (accessToken string, refreshToken string, err error) {

	// implementation later
	return
}
