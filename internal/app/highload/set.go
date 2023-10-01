package highload

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	desc "service-component/pkg/service-component/pb"
)

func (i *Implementation) Set(ctx context.Context, req *desc.SetRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, status.Error(codes.Unimplemented, "unimplemented")
}
