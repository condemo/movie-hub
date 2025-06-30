package types

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
)

type ServiceDataHandler interface {
	// TODO: debería devolver una lista de películas
	GetLastMovies(ctx context.Context) *pb.MediaResponse
}
