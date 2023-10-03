package main

import (
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"service-component/internal/app/highload"
	"service-component/pkg/service-component/pb"
)

func initGrpc(wg *sync.WaitGroup, service *highload.Implementation) {
	wg.Add(1)
	grpcServer := grpc.NewServer()
	pb.RegisterHighloadServiceServer(grpcServer, service)
	reflection.Register(grpcServer)
	lsn, err := net.Listen("tcp", ":7002")
	if err != nil {
		panic(err)
	}
	go func() {
		defer wg.Done()
		if err = grpcServer.Serve(lsn); err != nil {
			panic(err)
		}
	}()
}
