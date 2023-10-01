package highload

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "service-component/pkg/service-component/pb"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	return &desc.GetResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}
