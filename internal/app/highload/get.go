package highload

import (
	"context"
	"errors"

	"github.com/BornikReal/server-component/internal/app/infrastructure/logger"
	desc "github.com/BornikReal/server-component/pkg/service-component/pb"
	"github.com/BornikReal/storage-component/pkg/storage"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
