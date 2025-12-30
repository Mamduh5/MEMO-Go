package grpc

import (
	"context"

	"memo-go/services/auth/internal/usecase/auth"
	authv1 "memo-go/shared/gen/auth/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	authUC *auth.AuthUsecase
	authv1.UnimplementedAuthServiceServer
}

func NewAuthHandler(uc *auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC: uc}
}

func UserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDContextKey).(string)
	return userID, ok
}

func (h *AuthHandler) Login(
	ctx context.Context,
	req *authv1.LoginRequest,
) (*authv1.LoginResponse, error) {

	access, refresh, err := h.authUC.Login(
		ctx,
		req.Email,
		req.Password,
	)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &authv1.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (h *AuthHandler) Register(
	ctx context.Context,
	req *authv1.RegisterRequest,
) (*authv1.RegisterResponse, error) {

	err := h.authUC.Register(
		ctx,
		req.Email,
		req.Password,
	)
	if err != nil {
		switch err {
		case auth.ErrEmailAlreadyExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &authv1.RegisterResponse{}, nil
}

func (h *AuthHandler) Refresh(
	ctx context.Context,
	req *authv1.RefreshRequest,
) (*authv1.LoginResponse, error) {

	access, refresh, err := h.authUC.Refresh(ctx, req.RefreshToken)
	if err != nil {
		switch err {
		case auth.ErrInvalidRefreshToken:
			return nil, status.Error(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &authv1.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (h *AuthHandler) Logout(
	ctx context.Context,
	req *authv1.LogoutRequest,
) (*authv1.LogoutResponse, error) {

	userID, ok := UserIDFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing user context")
	}

	if err := h.authUC.Logout(ctx, userID); err != nil {
		return nil, err
	}

	return &authv1.LogoutResponse{}, nil
}
