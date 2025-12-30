package app

import (
	"net"

	"memo-go/services/auth/internal/config"
	authGrpc "memo-go/services/auth/internal/interface/transport/grpc"
	authv1 "memo-go/shared/gen/auth/v1"

	authgrpc "memo-go/services/auth/internal/interface/transport/grpc"
	"memo-go/services/auth/internal/usecase/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(
	cfg *config.Config,
	authUC *auth.AuthUsecase,
) error {

	lis, err := net.Listen("tcp", ":"+cfg.Server.GRPCPort)
	if err != nil {
		return err
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			authgrpc.JWTUnaryInterceptor([]byte(cfg.JWT.Secret)),
		),
	)

	reflection.Register(server)
	authv1.RegisterAuthServiceServer(
		server,
		authGrpc.NewAuthHandler(authUC),
	)

	return server.Serve(lis)
}
