package app

import (
	"net"

	authGrpc "memo-go/services/auth/internal/interface/transport/grpc"
	authv1 "memo-go/shared/gen/auth/v1"

	"memo-go/services/auth/internal/usecase/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(
	authUC *auth.AuthUsecase,
) error {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	reflection.Register(server)
	authv1.RegisterAuthServiceServer(
		server,
		authGrpc.NewAuthHandler(authUC),
	)

	return server.Serve(lis)
}
