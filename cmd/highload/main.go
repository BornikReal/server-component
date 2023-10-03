package main

import (
	"net/http"
	"sync"

	"github.com/BornikReal/storage-component/pkg/storage"
	"service-component/internal/app/highload"
	"service-component/internal/app/key_value"
)

func main() {
	storageService := storage.NewInMemoryStorage()
	kvRepository := key_value.NewRepository(storageService)
	kvService := key_value.NewKeyValueService(kvRepository)
	mainService := highload.NewImplementation(kvService)
	httpMux := initHttp(mainService)
	grpcServer, lsn := initGrpc(mainService)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		//err := http.ListenAndServe(os.Getenv("SERVERPORT"), httpMux)
		defer wg.Done()
		if err := http.ListenAndServe(":7000", httpMux); err != nil {
			panic(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcServer.Serve(lsn); err != nil {
			panic(err)
		}
	}()
	wg.Wait()
}
