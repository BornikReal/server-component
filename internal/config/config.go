package config

import (
	"fmt"
	"github.com/BornikReal/server-component/pkg/logger"
	"os"
	"strconv"
)

type Config struct {
	httpPort string
	grpcPort string

	ssDirectory string
	walPath     string
	walName     string

	maxTreeSize int
	blockSize   int64
	batch       int64

	compressCronJob string
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
