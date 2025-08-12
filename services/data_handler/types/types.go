package types

import (
	"context"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
)

type ServiceDataHandler interface {
	GetLastUpdates(context.Context, *pb.LastUpdatesRequest) (*pb.MediaListResponse, error)
	GetOneMedia(context.Context, int64) (*pb.Media, error)
	GetMediaFiltered(context.Context, *pb.MediaFilteredRequest) (*pb.MediaListResponse, error)
	DeleteMedia(context.Context, int64) error
	UpdateMedia(context.Context, *pb.Media) (*pb.Media, error)
	UpdateMediaBooleans(context.Context, *pb.MediaUpdateBool) (*pb.MediaResume, error)
}
