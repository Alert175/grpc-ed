package user_v1

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpc_user_v1 "grpc-ed/pkg/grpc/user_v1/api/grpc/user_v1"
	"log"
	"net"
)

type GrpcUserV1Server struct {
	grpc_user_v1.UnimplementedUserV1Server
}

// StartGrpcUserServer запустить сервер на порту
func StartGrpcUserServer() {
	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalf("fail start tcp listener, err - %s", err.Error())
	}
	server := grpc.NewServer()
	reflection.Register(server) // необходимо чтобы клиенты могли получать метаданные о сервере
	grpc_user_v1.RegisterUserV1Server(server, &GrpcUserV1Server{})
	log.Printf("grpc server listen at: %s", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("fail to serve: %s", err.Error())
	}
}
