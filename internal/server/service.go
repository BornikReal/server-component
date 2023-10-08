package server

import desc "github.com/BornikReal/server-component/pkg/service-component/pb"

type StorageService interface { //Storage
	Get(key string) (string, error)
	Set(key string, value string) error
}

type Implementation struct {
	desc.UnsafeHighloadServiceServer

	kvService StorageService
}

func NewImplementation(kvService StorageService) *Implementation {
	return &Implementation{
		kvService: kvService,
	}
}
