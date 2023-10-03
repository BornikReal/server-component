package main

import (
	"context"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"service-component/pkg/service-component/pb"
)

func initHttp(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	serveMux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterHighloadServiceHandlerFromEndpoint(ctx, serveMux, ":7002", opt)
	if err != nil {
		panic(err)
	}

	go func() {
		defer wg.Done()
		if err = http.ListenAndServe(":7000", serveMux); err != nil {
			panic(err)
		}
	}()
}
