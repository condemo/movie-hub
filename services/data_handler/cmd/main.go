package main

import (
	"github.com/condemo/movie-hub/services/common/config"
	datahandler "github.com/condemo/movie-hub/services/data_handler"
)

func main() {
	grpcServer := datahandler.NewGrpcServer(config.EnvConfig.DataGrpcPort)
	grpcServer.Run()
}
