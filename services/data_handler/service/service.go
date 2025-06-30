package service

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
)

type DataService struct {
	// Injections
}

func NewDataService() *DataService {
	return &DataService{}
}

func (s *DataService) GetLastMovies(ctx context.Context) *pb.MediaResponse {
	// TODO: Esto debería devolver una lista de pelis
	return &pb.MediaResponse{
		Msg: "this is a movie, promise",
	}
}
