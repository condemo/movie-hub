package utils

import (
	"log"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/common/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func FromPBMediaToTypeMedia(pb *pb.Media) *types.Media {
	return &types.Media{
		Id:          pb.GetId(),
		Type:        pb.GetType(),
		Title:       pb.GetTitle(),
		Year:        pb.GetYear(),
		FirstAir:    pb.GetFirstAir(),
		Genres:      pb.GetGenres(),
		Seasons:     pb.GetSeasons(),
		Caps:        pb.GetCaps(),
		Description: pb.GetDescription(),
		Rating:      pb.GetRating(),
		Image:       pb.GetImage(),
		Fav:         pb.GetFav(),
		Viewed:      pb.GetViewed(),
	}
}
