package handlers

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"google.golang.org/grpc"
)

type DataHandler struct {
	pb.UnimplementedDataHandlerServer
}

func NewDataHandler(grpc *grpc.Server) {
	grpcHandler := &DataHandler{}
	pb.RegisterDataHandlerServer(grpc, grpcHandler)
}

func (h *DataHandler) GetLastUpdates(ctx context.Context, lu *pb.LastUpdatesRequest) (*pb.MediaResponse, error) {
	// TODO:
	return &pb.MediaResponse{
		Msg: "This is a movie, promise",
	}, nil
}
