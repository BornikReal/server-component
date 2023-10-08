package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	httpPort string
	grpcPort string
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadFromEnv() error {
	httpPort := os.Getenv("SERVICE_PORT_HTTP")
	if httpPort == "" {
		return errors.New("SERVICE_PORT_HTTP is empty")
	}
	c.httpPort = fmt.Sprintf(":%s", httpPort)

	grpcPort := os.Getenv("SERVICE_PORT_GRPC")
	if grpcPort == "" {
		return errors.New("SERVICE_PORT_GRPC is empty")
	}
	c.grpcPort = fmt.Sprintf(":%s", grpcPort)
	return nil
}

func (c *Config) GetHttpPort() string {
	return c.httpPort
}

func (c *Config) GetGrpcPort() string {
	return c.grpcPort
}
