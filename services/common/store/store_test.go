package store

import (
	"os"
	"testing"

	_ "github.com/condemo/movie-hub/services/common/config"
	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/common/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockupDB *Storage

var mockupMovie types.Media = types.Media{
	// Id:          1,
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

var mockupMovie2 types.Media = types.Media{
	// Id:          1,
	Type:        "movie",
	Title:       "Fake Movie 2",
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

var mockupSerie pb.Media = pb.Media{
	// Id:          2,
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

var mockupResume types.MediaResume = types.MediaResume{
	// Id:          1,
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

func TestInsertBulkMedia(t *testing.T) {
	err := mockupDB.InsertBulkMedia([]types.Media{
		mockupMovie,
		mockupMovie2,
	})

	require.NoError(t, err)

	assert.NotZero(t, mockupMovie)
}
