package handlers

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/data_handler/types"
	"google.golang.org/grpc"
)

type DataHandler struct {
	pb.UnimplementedDataHandlerServer
	dataService types.ServiceDataHandler
}

func NewDataHandler(grpc *grpc.Server, s types.ServiceDataHandler) {
	grpcHandler := &DataHandler{dataService: s}
	pb.RegisterDataHandlerServer(grpc, grpcHandler)
}

func (h *DataHandler) GetLastUpdates(ctx context.Context, lu *pb.LastUpdatesRequest) (*pb.MediaResponse, error) {
	// TODO: implementar de verdad
	mr := h.dataService.GetLastMovies(ctx)
	return mr, nil
}
