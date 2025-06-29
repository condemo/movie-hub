package main

import (
	"fmt"

	datahandler "github.com/condemo/movie-hub/services/data_handler"
)

func main() {
	grpcServer := datahandler.NewGrpcServer(":5100")
	fmt.Println("DataGrpc Running on port: 5100")
	grpcServer.Run()
}
