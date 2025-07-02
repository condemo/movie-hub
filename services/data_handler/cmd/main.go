package main

import (
	"github.com/condemo/movie-hub/services/common/config"
	"github.com/condemo/movie-hub/services/common/store"
	datahandler "github.com/condemo/movie-hub/services/data_handler"
)

func main() {
	pqDB := store.NewPostgresqlStorage()
	// TODO:
	_ = store.NewStorage(pqDB)
	grpcServer := datahandler.NewGrpcServer(config.EnvConfig.DataGrpcPort)
	grpcServer.Run()
}
