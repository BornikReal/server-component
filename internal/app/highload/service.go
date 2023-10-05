package highload

import desc "github.com/BornikReal/server-component/pkg/service-component/pb"

type KeyValueService interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type Implementation struct {
	desc.UnsafeHighloadServiceServer

	kvService KeyValueService
}

func NewImplementation(kvService KeyValueService) *Implementation {
	return &Implementation{
		kvService: kvService,
	}
}
