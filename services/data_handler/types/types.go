package types

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
)

type ServiceDataHandler interface {
	// TODO: debería devolver una lista de películas
	GetLastUpdates(ctx context.Context) *pb.MediaListResponse
	GetOneMedia(ctx context.Context, id int64) *pb.Media
	GetMediaFiltered(ctx context.Context, fb *pb.FilterBy) *pb.MediaListResponse
}
