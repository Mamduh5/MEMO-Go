package app

import (
	"net"

	"memo-go/services/pos/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"memo-go/services/pos/internal/usecase/pos"

	posgrpc "memo-go/services/pos/internal/interface/transport/grpc"
	authv1 "memo-go/shared/gen/auth/v1"
	posv1 "memo-go/shared/gen/pos/v1"
)

func StartGRPCServer(
	cfg *config.Config,
	posUC *pos.PosUsecase,
) error {
	lis, err := net.Listen("tcp", ":"+cfg.Server.POSGRPCPort)
	if err != nil {
		return err
	}

	jwtInterceptor := posgrpc.NewJWTInterceptor(cfg.JWT.Secret)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(jwtInterceptor.Unary()),
	)

	conn, err := grpc.Dial(
		cfg.Server.AUTHGRPCPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	authClient := authv1.NewAuthServiceClient(conn)

	reflection.Register(server)
	posv1.RegisterPosServiceServer(
		server,
		posgrpc.NewPosHandler(authClient, posUC),
	)
	return server.Serve(lis)
}
