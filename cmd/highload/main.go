package main

import (
	"context"
	"sync"

	"github.com/BornikReal/storage-component/pkg/storage"
	"service-component/internal/app/highload"
	"service-component/internal/app/key_value"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	storageService := storage.NewInMemoryStorage()
	kvRepository := key_value.NewRepository(storageService)
	kvService := key_value.NewKeyValueService(kvRepository)
	mainService := highload.NewImplementation(kvService)

	wg := &sync.WaitGroup{}
	initHttp(ctx, wg)
	initGrpc(wg, mainService)
	wg.Wait()
}
