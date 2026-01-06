package grpc

import (
	"context"
	"log"
	"time"

	posuc "memo-go/services/pos/internal/usecase/pos"
	authv1 "memo-go/shared/gen/auth/v1"
	posv1 "memo-go/shared/gen/pos/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type PosHandler struct {
	posv1.UnimplementedPosServiceServer
	authClient authv1.AuthServiceClient
	uc         *posuc.Usecase
}

func NewPosHandler(authClient authv1.AuthServiceClient, uc *posuc.Usecase) *PosHandler {
	return &PosHandler{
		authClient: authClient,
		uc:         uc,
	}
}

func (h *PosHandler) Ping(
	ctx context.Context,
	_ *posv1.PingRequest,
) (*posv1.PingResponse, error) {

	userID, ok := UserIDFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing user context")
	}

	return &posv1.PingResponse{
		UserId:  userID,
		Message: "pong from POS",
	}, nil
}

func (h *PosHandler) Logout(
	ctx context.Context,
	_ *posv1.LogoutRequest,
) (*posv1.LogoutResponse, error) {

	// 1. Extract incoming metadata (JWT lives here)
	md, ok := metadata.FromIncomingContext(ctx)
	log.Println(md)
	if !ok {
		// no metadata = unauthenticated
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	// 2. Attach metadata to outgoing context
	outCtx := metadata.NewOutgoingContext(ctx, md)

	// 3. Call Auth with forwarded JWT
	_, err := h.authClient.Logout(outCtx, &authv1.LogoutRequest{})
	if err != nil {
		return nil, err
	}

	return &posv1.LogoutResponse{}, nil
}

func (h *PosHandler) OpenShift(
	ctx context.Context,
	_ *posv1.OpenShiftRequest,
) (*posv1.OpenShiftResponse, error) {

	userID, ok := UserIDFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing user context")
	}

	res, err := h.uc.OpenShift(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &posv1.OpenShiftResponse{
		ShiftId:  res.ShiftID,
		OpenedAt: res.OpenedAt.Format(time.RFC3339),
	}, nil
}
