package types

import "github.com/condemo/movie-hub/services/common/protogen/pb"

type Media struct {
	Id          int64  `db:"id"`
	Type        string `db:"media_type"`
	Title       string `db:"title"`
	Year        int32  `db:"release_year"`
	FirstAir    int32  `db:"first_air"`
	Genres      string `db:"genres"`
	Seasons     int32  `db:"seasons"`
	Caps        int32  `db:"caps"`
	Description string `db:"description"`
	Rating      int32  `db:"rating"`
	Runtime     int32  `db:"runtime"`
	Image       string `db:"image"`
	Fav         bool   `db:"fav"`
	Viewed      bool   `db:"viewed"`
}

func (m Media) GetProtoData() *pb.Media {
	return &pb.Media{
		Id:          m.Id,
		Type:        m.Type,
		Title:       m.Title,
		Year:        m.Year,
		Genres:      m.Genres,
		FirstAir:    m.FirstAir,
		Seasons:     m.Seasons,
		Caps:        m.Caps,
		Description: m.Description,
		Rating:      m.Rating,
		Runtime:     m.Runtime,
		Image:       m.Image,
		Fav:         m.Fav,
		Viewed:      m.Viewed,
	}
}

type MediaResume struct {
	Id          int64  `db:"id"`
	Type        string `db:"media_type"`
	Title       string `db:"title"`
	Genres      string `db:"genres"`
	Description string `db:"description"`
	Image       string `db:"image"`
	Rating      int32  `db:"rating"`
	Fav         bool   `db:"fav"`
	Viewed      bool   `db:"viewed"`
}

func (m MediaResume) GetProtoData() *pb.MediaResume {
	return &pb.MediaResume{
		Id:          m.Id,
		Type:        m.Type,
		Title:       m.Title,
		Genres:      m.Genres,
		Description: m.Description,
		Image:       m.Image,
		Rating:      m.Rating,
		Fav:         m.Fav,
		Viewed:      m.Viewed,
	}
}
