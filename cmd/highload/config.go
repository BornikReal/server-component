package main

import (
	"fmt"
	"os"
)

var (
	httpPort = fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT_HTTP"))
	grpcPort = fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT_GRPC"))
)
