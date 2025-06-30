package datahandler

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/condemo/movie-hub/services/data_handler/handlers"
	"github.com/condemo/movie-hub/services/data_handler/service"
	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
}

func NewGrpcServer(addr string) *grpcServer {
	return &grpcServer{addr: addr}
}

func (s *grpcServer) Run() {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal(err)
	}

	//SERVICE
	dt := service.NewDataService()

	// GRPC
	gServer := grpc.NewServer()
	handlers.NewDataHandler(gServer, dt)

	go func() {
		fmt.Println("DataGrpc Running on port", s.addr)
		log.Fatal(gServer.Serve(listener))
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-sigC

	gServer.GracefulStop()
}
