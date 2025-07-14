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

func (s *DataService) GetLastUpdates(ctx context.Context) (*pb.MediaListResponse, error) {
	// TODO: recibir desde el cliente un count como parametro de esta funcion y pasarlo
	// a `GetLastUpdates` para usarlo de limit en la DB
	data, err := s.store.GetLastUpdates(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.MediaResume, len(data))
	for i, d := range data {
		res[i] = d.GetProtoData()
	}

	return &pb.MediaListResponse{
		MediaList: res,
	}, nil
}

func (s *DataService) GetOneMedia(ctx context.Context, id int64) (*pb.Media, error) {
	// TODO:
	return nil, nil
}

func (s *DataService) GetMediaFiltered(ctx context.Context, fb *pb.FilterBy) (*pb.MediaListResponse, error) {
	// TODO:
	return nil, nil
}
