package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/condemo/movie-hub/services/common/config"
	"github.com/condemo/movie-hub/services/common/persistant"
	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/common/store"
	"github.com/condemo/movie-hub/services/common/utils"
)

type DataService struct {
	// Injections
	store           store.Store
	dFetcher        *dataFetcher
	nextUpdateTimer *time.Ticker
}

func NewDataService(s store.Store) *DataService {
	dt := &DataService{
		store:    s,
		dFetcher: newDataFetcher(),
	}

	return dt
}

func (s *DataService) Init() {
	// TODO: Load duration from config
	s.nextUpdateTimer = time.NewTicker(config.General.UpdateTimeInterval)

	s.updateData()

	go func() {
		for {
			<-s.nextUpdateTimer.C
			s.updateData()
		}
	}()
}

// FIX: gestionar errores
func (s *DataService) updateData() {
	data, err := s.dFetcher.GetLastUpdates(persistant.RequestData.LastMediaDate)
	if err != nil {
		log.Fatal(err)
	}

	var date int64
	if len(data.getShowList()) > 0 {
		err = s.store.InsertBulkMedia(context.Background(), data.getShowList())
		if err != nil {
			log.Fatal("DB ERROR -> ", err)
		}
		date = data.Changes[len(data.Changes)-1].TimeStamp
		fmt.Println("DATA UPDATED")
	} else {
		date = time.Now().Unix()
	}

	persistant.RequestData.LastMediaDate = &date
	err = persistant.RequestData.Save()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *DataService) GetLastUpdates(ctx context.Context, ml *pb.LastUpdatesRequest) (*pb.MediaListResponse, error) {
	data, err := s.store.GetLastUpdates(ctx, ml)
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

func (s *DataService) DeleteMedia(ctx context.Context, id int64) error {
	err := s.store.DeleteMedia(ctx, id)
	return err
}

func (s *DataService) UpdateMedia(ctx context.Context, m *pb.Media) (*pb.Media, error) {
	media := utils.FromPBMediaToTypeMedia(m)
	err := s.store.UpdateMedia(ctx, media)
	if err != nil {
		return nil, err
	}
	return media.GetProtoData(), nil
}

func (s *DataService) UpdateMediaBooleans(ctx context.Context, mb *pb.MediaUpdateBool) (*pb.MediaResume, error) {
	data, err := s.store.UpdateMediaBooleans(ctx, mb)
	if err != nil {
		return nil, err
	}

	return data.GetProtoData(), nil
}

func (s *DataService) GetMediaCount(ctx context.Context) (*pb.MediaCount, error) {
	count, err := s.store.GetMediaCount(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.MediaCount{Count: int64(count)}, nil
}
