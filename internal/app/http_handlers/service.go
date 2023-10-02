package http_handlers

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "service-component/pkg/service-component/pb"
)

type GrpcService interface {
	Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error)
	Set(ctx context.Context, req *desc.SetRequest) (*emptypb.Empty, error)
}

type HttpService struct {
	grpcService GrpcService
}

func NewHttpService(grpcService GrpcService) *HttpService {
	return &HttpService{
		grpcService: grpcService,
	}
}
