package main

import (
	"github.com/gorilla/mux"
	"service-component/internal/app/http_handlers"
)

func initHttp(server *http_handlers.HttpService) *mux.Router {
	httpMux := mux.NewRouter()
	//httpMux.StrictSlash(true)

	httpMux.HandleFunc("/kv", server.Get).Methods("GET")
	httpMux.HandleFunc("/task", server.Set).Methods("SET")

	// Настройка middleware для логирования и обработки ошибок panic.
	//httpMux.Use(func(h http.Handler) http.Handler {
	//	return handlers.LoggingHandler(os.Stdout, h)
	//})
	//httpMux.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

	//log.Fatal(http.ListenAndServe(os.Getenv("SERVERPORT"), httpMux))

	return httpMux
}
