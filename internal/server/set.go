package server

import (
	"context"
	"github.com/BornikReal/server-component/pkg/logger"
	"go.uber.org/zap"

	desc "github.com/BornikReal/server-component/pkg/service-component/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Set(_ context.Context, req *desc.SetRequest) (*emptypb.Empty, error) {
	err := i.kvService.Set(req.Key, req.Value)
	if err != nil {
		logger.Error("Set: error",
			zap.String("error", err.Error()),
			zap.String("key", req.Key), zap.String("value", req.Value))
		return nil, status.Errorf(codes.Internal, "Get: %v", err)
	}
	return &emptypb.Empty{}, nil
}
