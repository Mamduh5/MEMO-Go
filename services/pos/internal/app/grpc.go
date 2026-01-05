package app

import (
	"net"

	"memo-go/services/pos/internal/config"

	"google.golang.org/grpc"

	posgrpc "memo-go/services/pos/internal/interface/transport/grpc"
)

func startGRPCServer(cfg *config.Config) error {
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		return err
	}

	jwtInterceptor := posgrpc.NewJWTInterceptor(cfg.JWTSecret)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(jwtInterceptor.Unary()),
	)
	return server.Serve(lis)
}
