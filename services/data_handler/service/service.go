package service

import (
	"context"
	"log"
	"time"

	"github.com/condemo/movie-hub/services/common/persistant"
	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/common/store"
)

type DataService struct {
	// Injections
	store           store.Store
	dFetcher        *dataFetcher
	nextUpdateTimer *time.Timer
}

func NewDataService(s store.Store) *DataService {
	dt := &DataService{
		store:           s,
		dFetcher:        newDataFetcher(),
		nextUpdateTimer: time.NewTimer(time.Minute),
	}

	// FIX: llamar al acabar el timer
	dt.updateData()

	return dt
}

// FIX: gestionar errores
func (s *DataService) updateData() {
	data, err := s.dFetcher.GetLastUpdates(persistant.RequestData.LastMediaDate)
	if err != nil {
		log.Fatal(err)
	}

	err = s.store.InsertBulkMedia(context.Background(), data.getShowList())
	if err != nil {
		log.Fatal("DB ERROR ->", err)
	}

	date := data.Changes[len(data.Changes)-1].TimeStamp
	persistant.RequestData.LastMediaDate = &date
	err = persistant.RequestData.Save()
	if err != nil {
		log.Fatal(err)
	}
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

func (s *DataService) GetMediaFiltered(ctx context.Context, fb *pb.MediaFilteredRequest) (*pb.MediaListResponse, error) {
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
