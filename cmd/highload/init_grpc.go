package main

import (
	"google.golang.org/grpc/reflection"
	"net"

	"google.golang.org/grpc"
	"service-component/internal/app/highload"
	"service-component/pkg/service-component/pb"
)

func initGrpc(service *highload.Implementation) (*grpc.Server, net.Listener) {
	grpcServer := grpc.NewServer()
	pb.RegisterServiceComponentServer(grpcServer, service)
	reflection.Register(grpcServer)
	lsn, err := net.Listen("tcp", ":7002")
	if err != nil {
		panic(err)
	}
	return grpcServer, lsn
}
