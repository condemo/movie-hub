package service

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/common/store"
)

type DataService struct {
	// Injections
	store store.Store
}

func NewDataService(s store.Store) *DataService {
	return &DataService{store: s}
}

func (s *DataService) GetLastMovies(ctx context.Context) *pb.MediaResponse {
	// TODO: Esto debería devolver una lista de pelis
	return &pb.MediaResponse{
		Msg: "this is a movie, promise",
	}
}
