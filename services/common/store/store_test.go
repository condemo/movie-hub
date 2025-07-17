package store

import (
	"context"
	"os"
	"testing"

	_ "github.com/condemo/movie-hub/services/common/config"
	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/common/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockupDB *Storage

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var mockupMovie types.Media = types.Media{
	Type:        "movie",
	Title:       "Test Movie",
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

var mockupMovie2 types.Media = types.Media{
	Type:        "movie",
	Title:       "Test Movie 2",
	Year:        1955,
	Genres:      "acción,fantasía",
	Seasons:     0,
	Caps:        0,
	Description: "Some desc",
	Rating:      90,
	Image:       "image.jpg",
	Fav:         true,
	Viewed:      false,
}

var mockupSerie types.Media = types.Media{
	Type:        "serie",
	Title:       "Test Serie",
	Year:        1993,
	Genres:      "acción,fantasía",
	Seasons:     2,
	Caps:        21,
	Description: "Some desc",
	Rating:      91,
	Image:       "image.jpg",
	Fav:         false,
	Viewed:      true,
}

var mockupResume types.MediaResume = types.MediaResume{
	Type:        "movie",
	Title:       "Fake Movie",
	Genres:      "acción,fantasía",
	Description: "Some desc",
	Image:       "image.jpg",
	Fav:         true,
	Viewed:      false,
}

func TestMain(m *testing.M) {
	db := NewPostgresqlStorage()
	mockupDB = NewStorage(db)

	os.Exit(m.Run())
}

func TestInsertMedia(t *testing.T) {
	err := mockupDB.InsertMedia(context.Background(), &mockupMovie)
	require.NoError(t, err)

	err = mockupDB.InsertMedia(context.Background(), &mockupSerie)
	require.NoError(t, err)
}

func TestGetOneMedia(t *testing.T) {
	data, err := mockupDB.GetOneMedia(context.Background(), mockupMovie.Id)
	require.NoError(t, err)
	assert.Equal(t, &mockupMovie, data)
}

func TestUpdateMedia(t *testing.T) {
	mockupMovie.Title = "Updated Title"
	err := mockupDB.UpdateMedia(context.Background(), &mockupMovie)
	require.NoError(t, err)
}

func TestMediaFiltered(t *testing.T) {
	media, err := mockupDB.GetMediaFiltered(context.Background(), &pb.MediaFilteredRequest{
		Filter: pb.FilterBy_fav,
	})
	require.NoError(t, err)
	assert.Equal(t, mockupMovie.Fav, media[0].Fav)

	media, err = mockupDB.GetMediaFiltered(context.Background(), &pb.MediaFilteredRequest{
		Filter: pb.FilterBy_viewed,
	})
	require.NoError(t, err)
	assert.Equal(t, mockupSerie.Viewed, media[0].Viewed)
}

func TestDeleteMedia(t *testing.T) {
	err := mockupDB.DeleteMedia(context.Background(), mockupMovie.Id)
	require.NoError(t, err)

	err = mockupDB.DeleteMedia(context.Background(), mockupSerie.Id)
	require.NoError(t, err)
}
