package highload

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	desc "service-component/pkg/service-component/pb"
)

func (i *Implementation) Set(_ context.Context, req *desc.SetRequest) (*emptypb.Empty, error) {
	err := i.kvService.Set(req.Key, req.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Get: %v", err)
	}
	return &emptypb.Empty{}, nil
}
