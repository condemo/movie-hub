package service

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/common/store"
)

var mockupMovie pb.Media = pb.Media{
	Id:          1,
	Type:        "movie",
	Title:       "Fake Movie",
	Year:        1953,
	Genres:      "acción,fantasía",
	Seasons:     0,
	Caps:        0,
	Description: "Some desc",
	Rating:      78,
	Image:       "image.jpg",
	Fav:         true,
	Viewed:      false,
}

var mockupSerie pb.Media = pb.Media{
	Id:          2,
	Type:        "serie",
	Title:       "Fake Serie",
	Year:        1993,
	Genres:      "acción,fantasía",
	Seasons:     2,
	Caps:        21,
	Description: "Some desc",
	Rating:      91,
	Image:       "image.jpg",
	Fav:         false,
	Viewed:      false,
}

var mockupResume pb.MediaResume = pb.MediaResume{
	Id:          1,
	Type:        "movie",
	Title:       "Fake Movie",
	Genres:      "acción,fantasía",
	Description: "Some desc",
	Image:       "image.jpg",
	Fav:         true,
	Viewed:      false,
}

type DataService struct {
	// Injections
	store store.Store
}

func NewDataService(s store.Store) *DataService {
	return &DataService{store: s}
}

func (s *DataService) GetLastMovies(ctx context.Context) *pb.MediaListResponse {
	// TODO: Esto debería devolver una lista de pelis
	return &pb.MediaListResponse{
		MediaList: []*pb.MediaResume{
			&mockupResume,
			&mockupResume,
			&mockupResume,
		},
	}
}

func (s *DataService) GetMovie(ctx context.Context, id int64) *pb.Media {
	// TODO:
	return &mockupMovie
}

func (s *DataService) GetSerie(ctx context.Context, id int64) *pb.Media {
	// TODO:
	return &mockupSerie
}
