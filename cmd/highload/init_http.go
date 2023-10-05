package main

import (
	"context"
	"net/http"
	"sync"

	"github.com/BornikReal/server-component/internal/app/infrastructure/logger"
	"github.com/BornikReal/server-component/pkg/service-component/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initHttp(ctx context.Context, wg *sync.WaitGroup) {
	logger.Info("init http server")
	wg.Add(1)
	serveMux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterHighloadServiceHandlerFromEndpoint(ctx, serveMux, grpcPort, opt)
	if err != nil {
		logger.Fatal("can't create http server from grpc endpoint",
			zap.String("error", err.Error()))
	}

	go func() {
		defer wg.Done()
		if err = http.ListenAndServe(httpPort, serveMux); err != nil {
			logger.Fatal("starting http server ended with error",
				zap.String("error", err.Error()), zap.String("port", httpPort))
		}
	}()
}
