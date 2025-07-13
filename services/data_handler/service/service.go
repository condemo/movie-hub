package service

import (
	"context"
	"log"

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

func (s *DataService) GetLastUpdates(ctx context.Context) *pb.MediaListResponse {
	// TODO: recibir desde el cliente un count como parametro de esta funcion y pasarlo
	// a `GetLastUpdates` para usarlo de limit en la DB
	data, err := s.store.GetLastUpdates(ctx)
	if err != nil {
		// TODO: El servicio debería devolver errores en métodos que se utilicen en el handler grpc
		log.Fatal(err)
	}

	res := make([]*pb.MediaResume, len(data))
	for i, d := range data {
		res[i] = d.GetProtoData()
	}

	return &pb.MediaListResponse{
		MediaList: res,
	}
}

func (s *DataService) GetOneMedia(ctx context.Context, id int64) *pb.Media {
	// TODO:
	return nil
}

func (s *DataService) GetMediaFiltered(ctx context.Context, fb *pb.FilterBy) *pb.MediaListResponse {
	// TODO:
	return nil
}
