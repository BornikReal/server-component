package main

import (
	"github.com/gorilla/mux"
	"service-component/internal/app/highload"
	"service-component/internal/app/http_handlers"
)

func initHttp(service *highload.Implementation) *mux.Router {
	httpService := http_handlers.NewHttpService(service)
	httpMux := mux.NewRouter()
	//httpMux.StrictSlash(true)

	httpMux.HandleFunc("/kv", httpService.Get).Methods("GET")
	httpMux.HandleFunc("/kv", httpService.Set).Methods("POST")

	// Настройка middleware для логирования и обработки ошибок panic.
	//httpMux.Use(func(h http.Handler) http.Handler {
	//	return handlers.LoggingHandler(os.Stdout, h)
	//})
	//httpMux.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

	//log.Fatal(http.ListenAndServe(os.Getenv("SERVERPORT"), httpMux))

	return httpMux
}
