package internal

import (
	"context"
	"github.com/BornikReal/server-component/internal/config"
	"github.com/BornikReal/server-component/internal/server"
	"github.com/BornikReal/server-component/internal/storage_service/inmemory"
	"github.com/BornikReal/server-component/pkg/logger"
	"github.com/BornikReal/server-component/pkg/service-component/pb"
	"github.com/BornikReal/storage-component/pkg/ss_storage/iterator"
	"github.com/BornikReal/storage-component/pkg/ss_storage/kv_file"
	"github.com/BornikReal/storage-component/pkg/ss_storage/ss_manager"
	"github.com/BornikReal/storage-component/pkg/storage"
	"github.com/BornikReal/storage-component/pkg/tree_with_clone"
	"github.com/emirpasic/gods/trees/avltree"
	"github.com/go-co-op/gocron"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
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

	conf := config.New()

	if err := conf.LoadFromEnv(); err != nil {
		return err
	}

	var mt inmemory.Storage
	switch conf.GetStorageType() {
	case config.LSMStorage:
		mt = initLSMStorage(conf)
	case config.RedisClusterStorage:
		mt = initRedisCluster(conf)
	}

	if mt == nil {
		logger.Fatalf("storage not init",
			zap.String("storage_type", string(conf.GetStorageType())))
	}

	storageService := inmemory.NewStorageService(mt)
	impl := server.NewImplementation(storageService)

	app.config = conf

	app.initGrpc(impl)
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

func initLSMStorage(conf *config.Config) inmemory.Storage {
	ssManager := ss_manager.NewSSManager(conf.GetSSDirectory(), conf.GetBlockSize(), conf.GetBatch())
	if err := ssManager.Init(); err != nil {
		panic(err)
	}
	tree := avltree.NewWithStringComparator()
	wal := kv_file.NewKVFile(conf.GetWalPath(), conf.GetWalName())
	if err := wal.Init(); err != nil {
		panic(err)
	}

	dumper := make(chan iterator.Iterator, conf.SSChanSize())
	mt := storage.NewMemTableWithWal(
		storage.NewMemTableWithSS(
			storage.NewMemTable(
				tree_with_clone.NewTreeWithClone(
					tree,
				),
				dumper,
				conf.GetMaxTreeSize(),
			),
			ssManager,
		),
		wal,
	)

	errorCh := make(chan error, 1)
	ssProcessor := storage.NewSSProcessor(ssManager, errorCh)
	go ssProcessor.Start(dumper)
	go func() {
		for err := range errorCh {
			logger.Error("SS processor encounters with error while saving tree", zap.String("error", err.Error()))
		}
	}()

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Cron(conf.GetCompressCronJob()).Do(ssManager.CompressSS)
	if err != nil {
		panic(err)
	}
	s.StartAsync()

	return mt
}

func initRedisReplication(conf *config.Config) inmemory.Storage {
	rdbMaster := redis.NewClient(&redis.Options{
		Addr:     conf.GetMasterHost(),
		Password: conf.GetMasterRedisPassword(),
	})

	rdbSlave1 := redis.NewClient(&redis.Options{
		Addr:     conf.GetSlave1RedisHost(),
		Password: conf.GetSlave1RedisPassword(),
	})

	rdbSlave2 := redis.NewClient(&redis.Options{
		Addr:     conf.GetSlave2RedisHost(),
		Password: conf.GetSlave2RedisPassword(),
	})

	return storage.NewRedisStorage(rdbMaster, []storage.RedisClient{rdbSlave1, rdbSlave2})
}

func initRedisCluster(conf *config.Config) inmemory.Storage {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(conf.GetClusterRedisHosts(), ","),
		Password: conf.GetClusterRedisPassword(),
	})

	return storage.NewRedisStorage(rdb, nil)
}
