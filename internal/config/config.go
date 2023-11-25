package config

import (
	"fmt"
	"github.com/BornikReal/server-component/pkg/logger"
	"os"
	"strconv"
)

type StorageType uint8

const (
	LSMStorage          = 0
	RedisClusterStorage = 1
)

var storageTypeConv = map[string]StorageType{
	"lsm":           LSMStorage,
	"redis_cluster": RedisClusterStorage,
}

type Config struct {
	httpPort string
	grpcPort string

	ssDirectory string
	walPath     string
	walName     string

	maxTreeSize int
	blockSize   int64
	batch       int64
	ssChanSize  int64

	compressCronJob string

	masterRedisHost     string
	masterRedisPassword string

	slave1RedisHost     string
	slave1RedisPassword string

	slave2RedisHost     string
	slave2RedisPassword string

	storageType StorageType

	clusterRedisHosts    string
	clusterRedisPassword string
}

func New() *Config {
	return &Config{}
}

func logUseDefault(varName string, stdValue interface{}) {
	logger.Infof("%s not found, using standard value: %v", varName, stdValue)
}

func (c *Config) LoadFromEnv() error {
	value := os.Getenv("SERVICE_PORT_HTTP")
	if value == "" {
		value = "7001"
		logUseDefault("SERVICE_PORT_HTTP", value)
	}
	c.httpPort = fmt.Sprintf(":%s", value)

	value = os.Getenv("SERVICE_PORT_GRPC")
	if value == "" {
		value = "7002"
		logUseDefault("SERVICE_PORT_GRPC", value)
	}
	c.grpcPort = fmt.Sprintf(":%s", value)

	c.ssDirectory = os.Getenv("DB_DIR")
	if c.ssDirectory == "" {
		c.ssDirectory = "db"
		logUseDefault("DB_DIR", c.ssDirectory)
	}

	c.compressCronJob = os.Getenv("COMPRESS_CRON_JOB")
	if c.compressCronJob == "" {
		c.compressCronJob = "0 */1 * * *"
		logUseDefault("COMPRESS_CRON_JOB", c.compressCronJob)
	}

	value = os.Getenv("MAX_TREE_SIZE")
	if value == "" {
		c.maxTreeSize = 5
		logUseDefault("MAX_TREE_SIZE", c.maxTreeSize)
	} else {
		valueInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			c.maxTreeSize = 5
			logUseDefault("MAX_TREE_SIZE", c.maxTreeSize)
		} else {
			c.maxTreeSize = int(valueInt)
		}
	}

	value = os.Getenv("BLOCK_SIZE")
	if value == "" {
		c.blockSize = 5
		logUseDefault("BLOCK_SIZE", c.blockSize)
	} else {
		valueInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			c.blockSize = 5
			logUseDefault("BLOCK_SIZE", c.blockSize)
		} else {
			c.blockSize = valueInt
		}
	}

	value = os.Getenv("BATCH")
	if value == "" {
		c.batch = 1
		logUseDefault("BATCH", c.batch)
	} else {
		valueInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			c.batch = 1
			logUseDefault("BATCH", c.batch)
		} else {
			c.batch = valueInt
		}
	}

	c.walPath = os.Getenv("WAL_PATH")
	if c.walPath == "" {
		c.walPath = ""
		logUseDefault("WAL_PATH", c.walPath)
	}

	c.walName = os.Getenv("WAL_NAME")
	if c.walName == "" {
		c.walName = ""
		logUseDefault("WAL_NAME", c.walName)
	}

	value = os.Getenv("SS_CHAN_SIZE")
	if value == "" {
		c.ssChanSize = 5
		logUseDefault("SS_CHAN_SIZE", c.ssChanSize)
	} else {
		valueInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			c.ssChanSize = 5
			logUseDefault("SS_CHAN_SIZE", c.ssChanSize)
		} else {
			c.ssChanSize = valueInt
		}
	}

	c.masterRedisHost = os.Getenv("MASTER_REDIS_HOST")
	if c.masterRedisHost == "" {
		c.masterRedisHost = "172.28.1.4:6380"
		logUseDefault("MASTER_REDIS_HOST", c.masterRedisHost)
	}

	c.masterRedisPassword = os.Getenv("MASTER_REDIS_PASSWORD")
	if c.masterRedisPassword == "" {
		c.masterRedisPassword = "1234"
		logUseDefault("MASTER_REDIS_PASSWORD", c.masterRedisPassword)
	}

	c.slave1RedisHost = os.Getenv("SLAVE1_REDIS_HOST")
	if c.slave1RedisHost == "" {
		c.slave1RedisHost = "172.28.1.5:6381"
		logUseDefault("SLAVE1_REDIS_HOST", c.slave1RedisHost)
	}

	c.slave1RedisPassword = os.Getenv("SLAVE1_REDIS_PASSWORD")
	if c.slave1RedisPassword == "" {
		c.slave1RedisPassword = "1234"
		logUseDefault("SLAVE1_REDIS_PASSWORD", c.slave1RedisPassword)
	}

	c.slave2RedisHost = os.Getenv("SLAVE1_REDIS_HOST")
	if c.slave2RedisHost == "" {
		c.slave2RedisHost = "172.28.1.6:6382"
		logUseDefault("SLAVE1_REDIS_HOST", c.slave2RedisHost)
	}

	c.slave2RedisPassword = os.Getenv("SLAVE2_REDIS_PASSWORD")
	if c.slave2RedisPassword == "" {
		c.slave2RedisPassword = "1234"
		logUseDefault("SLAVE2_REDIS_PASSWORD", c.slave2RedisPassword)
	}

	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "" {
		c.storageType = LSMStorage
		logUseDefault("STORAGE_TYPE", c.storageType)
	} else {
		var ok bool
		c.storageType, ok = storageTypeConv[storageType]
		if !ok {
			c.storageType = LSMStorage
			logger.Infof("Unknown storage type - %s, using default storage type - lsm", storageType)
		}
	}

	c.clusterRedisHosts = os.Getenv("CLUSTER_REDIS_HOSTS")
	if c.clusterRedisHosts == "" {
		c.clusterRedisHosts = "172.28.1.10:7000,172.28.1.11:7001,172.28.1.12:7002,172.28.1.13:7003,172.28.1.14:7004,172.28.1.15:7005"
		logUseDefault("CLUSTER_REDIS_HOSTS", c.clusterRedisHosts)
	}

	c.clusterRedisPassword = os.Getenv("CLUSTER_REDIS_PASSWORD")
	if c.clusterRedisPassword == "" {
		c.clusterRedisPassword = "1234"
		logUseDefault("CLUSTER_REDIS_PASSWORD", c.clusterRedisPassword)
	}

	return nil
}

func (c *Config) GetHttpPort() string {
	return c.httpPort
}

func (c *Config) GetGrpcPort() string {
	return c.grpcPort
}

func (c *Config) GetSSDirectory() string {
	return c.ssDirectory
}

func (c *Config) GetWalPath() string {
	return c.walPath
}

func (c *Config) GetWalName() string {
	return c.walName
}

func (c *Config) GetCompressCronJob() string {
	return c.compressCronJob
}

func (c *Config) GetMaxTreeSize() int {
	return c.maxTreeSize
}

func (c *Config) GetBlockSize() int64 {
	return c.blockSize
}

func (c *Config) GetBatch() int64 {
	return c.batch
}

func (c *Config) SSChanSize() int64 {
	return c.ssChanSize
}

func (c *Config) GetMasterHost() string {
	return c.masterRedisHost
}

func (c *Config) GetMasterRedisPassword() string {
	return c.masterRedisPassword
}

func (c *Config) GetSlave1RedisHost() string {
	return c.slave1RedisHost
}

func (c *Config) GetSlave1RedisPassword() string {
	return c.slave1RedisPassword
}

func (c *Config) GetSlave2RedisHost() string {
	return c.slave2RedisHost
}

func (c *Config) GetSlave2RedisPassword() string {
	return c.slave2RedisPassword
}

func (c *Config) GetStorageType() StorageType {
	return c.storageType
}

func (c *Config) GetClusterRedisHosts() string {
	return c.clusterRedisHosts
}

func (c *Config) GetClusterRedisPassword() string {
	return c.clusterRedisPassword
}
