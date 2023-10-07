package internal

import (
	"context"
	"net"
	"net/http"
	"sync"

	"github.com/BornikReal/server-component/internal/config"
	"github.com/BornikReal/server-component/internal/infrastructure/logger"
	"github.com/BornikReal/server-component/internal/server"
	"github.com/BornikReal/server-component/internal/storage_service"
	"github.com/BornikReal/server-component/pkg/service-component/pb"
	"github.com/BornikReal/storage-component/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type serve func() error

type App struct {
	wg        sync.WaitGroup
	config    *config.Config
	startHttp serve
	startGrpc serve
}

func NewApp() *App {
	return &App{}
}

func (app *App) Init() error {
	logger.InitLogger()
	logger.Info("init service")
	defer logger.Info("init finished")

	ctx := context.Background()
	storageService := storage.NewInMemoryStorage()
	kvRepository := storage_service.NewRepository(storageService)
	kvService := storage_service.NewKeyValueService(kvRepository)
	mainService := server.NewImplementation(kvService)

	conf := config.New()

	if err := conf.Init(); err != nil {
		return err
	}
	app.config = conf

	app.initGrpc(mainService)
	app.initHttp(ctx)

	return nil
}

func (app *App) Run() {
	logger.Info("service is starting")
	defer logger.Info("service shutdown")
	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		if err := app.startGrpc(); err != nil {
			logger.Fatal("starting grpc storage_service ended with error",
				zap.String("error", err.Error()), zap.String("port", app.config.GetGrpcPort()))
		}
	}()

	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		if err := app.startHttp(); err != nil {
			logger.Fatal("starting http storage_service ended with error",
				zap.String("error", err.Error()), zap.String("port", app.config.GetHttpPort()))
		}
	}()

	logger.Infof("Service successfully started. Ports: HTTP - %s, GRPC - %s",
		app.config.GetHttpPort()[1:], app.config.GetGrpcPort()[1:])
	app.wg.Wait()
}

func (app *App) initGrpc(service *server.Implementation) {
	logger.Info("init grpc storage_service")
	grpcServer := grpc.NewServer()
	pb.RegisterHighloadServiceServer(grpcServer, service)
	reflection.Register(grpcServer)
	lsn, err := net.Listen("tcp", app.config.GetGrpcPort())
	if err != nil {
		logger.Fatal("listening port ended with error",
			zap.String("error", err.Error()), zap.String("port", app.config.GetGrpcPort()))
	}

	app.startGrpc = func() error {
		return grpcServer.Serve(lsn)
	}
}

func (app *App) initHttp(ctx context.Context) {
	logger.Info("init http storage_service")
	serveMux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterHighloadServiceHandlerFromEndpoint(ctx, serveMux, app.config.GetGrpcPort(), opt)
	if err != nil {
		logger.Fatal("can't create http storage_service from grpc endpoint",
			zap.String("error", err.Error()))
	}

	app.startHttp = func() error {
		return http.ListenAndServe(app.config.GetHttpPort(), serveMux)
	}
}
