package main

import (
	"net"
	"sync"

	"github.com/BornikReal/server-component/internal/app/highload"
	"github.com/BornikReal/server-component/internal/app/infrastructure/logger"
	"github.com/BornikReal/server-component/pkg/service-component/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func initGrpc(wg *sync.WaitGroup, service *highload.Implementation) {
	logger.Info("init grpc server")
	wg.Add(1)
	grpcServer := grpc.NewServer()
	pb.RegisterHighloadServiceServer(grpcServer, service)
	reflection.Register(grpcServer)
	lsn, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.Fatal("listening port ended with error",
			zap.String("error", err.Error()), zap.String("port", grpcPort))
	}
	go func() {
		defer wg.Done()
		if err = grpcServer.Serve(lsn); err != nil {
			logger.Fatal("starting grpc server ended with error",
				zap.String("error", err.Error()), zap.String("port", grpcPort))
		}
	}()
}
