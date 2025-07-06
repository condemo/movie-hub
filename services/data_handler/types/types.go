package types

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
)

type ServiceDataHandler interface {
	// TODO: debería devolver una lista de películas
	GetLastMovies(ctx context.Context) *pb.MediaListResponse
	GetMovie(ctx context.Context, id int64) *pb.MediaResponse
	GetSerie(ctx context.Context, id int64) *pb.MediaResponse
}
