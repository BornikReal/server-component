package main

import (
	"github.com/BornikReal/storage-component/pkg/storage"
	"net/http"
	"os"
	"service-component/internal/app/highload"
	"service-component/internal/app/http_handlers"
	"service-component/internal/app/key_value"
)

func main() {
	storageService := storage.NewInMemoryStorage()
	kvRepository := key_value.NewRepository(storageService)
	kvService := key_value.NewKeyValueService(kvRepository)
	grpcService := highload.NewImplementation(kvService)
	httpService := http_handlers.NewHttpService(grpcService)
	httpMux := initHttp(httpService)
	err := http.ListenAndServe(os.Getenv("SERVERPORT"), httpMux)
	if err != nil {
		panic(err)
	}
}
