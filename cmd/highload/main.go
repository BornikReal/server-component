package main

import (
	"context"
	"sync"

	"github.com/BornikReal/server-component/internal/app/highload"
	"github.com/BornikReal/server-component/internal/app/infrastructure/logger"
	"github.com/BornikReal/server-component/internal/app/key_value"
	"github.com/BornikReal/storage-component/pkg/storage"
)

func main() {
	initLogger()
	logger.Info("init service")
	defer logger.Info("service shutdown")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	storageService := storage.NewInMemoryStorage()
	kvRepository := key_value.NewRepository(storageService)
	kvService := key_value.NewKeyValueService(kvRepository)
	mainService := highload.NewImplementation(kvService)

	wg := &sync.WaitGroup{}
	initGrpc(wg, mainService)
	initHttp(ctx, wg)
	logger.Infof("Service successfully started. Ports: HTTP - %s, GRPC - %s", httpPort[1:], grpcPort[1:])
	wg.Wait()
}
