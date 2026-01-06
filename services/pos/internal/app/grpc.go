package app

import (
	"net"

	"memo-go/services/pos/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	posgrpc "memo-go/services/pos/internal/interface/transport/grpc"
	posuc "memo-go/services/pos/internal/usecase/pos"
	authv1 "memo-go/shared/gen/auth/v1"
	posv1 "memo-go/shared/gen/pos/v1"
)

func startGRPCServer(cfg *config.Config) error {
	lis, err := net.Listen("tcp", ":"+cfg.POSGRPCPort)
	if err != nil {
		return err
	}

	jwtInterceptor := posgrpc.NewJWTInterceptor(cfg.JWTSecret)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(jwtInterceptor.Unary()),
	)

	reflection.Register(server)

	conn, err := grpc.Dial(
		cfg.AuthGRPCAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	authClient := authv1.NewAuthServiceClient(conn)
	uc := posuc.New()

	handler := posgrpc.NewPosHandler(authClient, uc)
	posv1.RegisterPosServiceServer(server, handler)
	return server.Serve(lis)
}
