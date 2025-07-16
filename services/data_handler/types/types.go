package types

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
)

type ServiceDataHandler interface {
	GetLastUpdates(ctx context.Context, limit int32) (*pb.MediaListResponse, error)
	GetOneMedia(ctx context.Context, id int64) (*pb.Media, error)
	GetMediaFiltered(ctx context.Context, fb pb.FilterBy) (*pb.MediaListResponse, error)
}
