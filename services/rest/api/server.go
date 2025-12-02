package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/condemo/movie-hub/services/common/config"
	"github.com/condemo/movie-hub/services/common/utils"
	"github.com/condemo/movie-hub/services/rest/api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr: addr,
	}
}

func (s *ApiServer) Run() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://192.168.3.*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}), middleware.Logger)

	server := http.Server{
		Addr:         s.addr,
		Handler:      r,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
	}

	// GRPC
	dataGrpc := utils.NewGrpcClient(config.EnvConfig.DataGrpcPort)

	mediaHandler := handlers.NewMediaHandler(dataGrpc)
	r.Mount("/movie", mediaHandler.RegisterRoutes())

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-sigC

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// server.Shutdown ends the execution of the program
	// after waiting for all active connections to finish or 30 seconds to pass
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Bad Shutdown ->", err)
	}
}
