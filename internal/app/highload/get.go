package highload

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"service-component/internal/app/infrastructure/logger"

	"github.com/BornikReal/storage-component/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "service-component/pkg/service-component/pb"
)

func (i *Implementation) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	value, err := i.kvService.Get(req.Key)
	if errors.Is(err, storage.NotFoundError) {
		logger.Info("Get: key not found", zap.String("key", req.Key))
		return nil, status.Errorf(codes.NotFound, "Get: key %s not found", req.Key)
	} else if err != nil {
		logger.Error("Get: error", zap.String("error", err.Error()), zap.String("key", req.Key))
		return nil, status.Errorf(codes.Internal, "Get: %v", err)
	}
	return &desc.GetResponse{
		Value: value,
	}, nil
}
