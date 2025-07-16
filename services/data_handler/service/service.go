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

func (s *DataService) GetLastUpdates(ctx context.Context, limit int32) (*pb.MediaListResponse, error) {
	data, err := s.store.GetLastUpdates(ctx, limit)
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
	media, err := s.store.GetOneMedia(ctx, id)
	if err != nil {
		return nil, err
	}

	return media.GetProtoData(), nil
}

func (s *DataService) GetMediaFiltered(ctx context.Context, fb pb.FilterBy) (*pb.MediaListResponse, error) {
	mediaFiltered, err := s.store.GetMediaFiltered(ctx, fb)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.MediaResume, len(mediaFiltered))
	for i, d := range mediaFiltered {
		res[i] = d.GetProtoData()
	}

	return &pb.MediaListResponse{MediaList: res}, nil
}
