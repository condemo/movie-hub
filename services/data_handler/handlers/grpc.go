package handlers

import (
	"context"
	"net/http"

	"github.com/condemo/movie-hub/services/common/errs"
	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/data_handler/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DataHandler struct {
	pb.UnimplementedDataHandlerServer
	dataService types.ServiceDataHandler
}

func NewDataHandler(grpc *grpc.Server, s types.ServiceDataHandler) {
	grpcHandler := &DataHandler{dataService: s}
	pb.RegisterDataHandlerServer(grpc, grpcHandler)
}

func (h *DataHandler) GetLastUpdates(ctx context.Context, lu *pb.LastUpdatesRequest) (*pb.MediaListResponse, error) {
	mr, err := h.dataService.GetLastUpdates(ctx, lu.GetLimit())
	if err != nil {
		return nil, err
	}
	return mr, nil
}

func (h *DataHandler) GetOneMedia(ctx context.Context, sr *pb.MediaRequest) (*pb.Media, error) {
	resp, err := h.dataService.GetOneMedia(ctx, sr.GetId())
	if err != nil {
		return nil, errs.NewApiError(err, http.StatusNotFound, "media not found")
	}

	return resp, nil
}

func (h *DataHandler) GetMediaFiltered(ctx context.Context, mr *pb.MediaFilteredRequest) (*pb.MediaListResponse, error) {
	res, err := h.dataService.GetMediaFiltered(ctx, mr)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *DataHandler) DeleteMedia(ctx context.Context, mr *pb.MediaRequest) (*emptypb.Empty, error) {
	err := h.dataService.DeleteMedia(ctx, mr.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *DataHandler) UpdateMedia(ctx context.Context, ur *pb.UpdateMediaReq) (*pb.MediaResponse, error) {
	res, err := h.dataService.UpdateMedia(ctx, ur.GetMedia())
	if err != nil {
		return nil, err
	}
	return &pb.MediaResponse{Media: res}, nil
}

func (h *DataHandler) UpdateMediaBooleans(ctx context.Context, mb *pb.MediaUpdateBool) (*pb.MediaResume, error) {
	res, err := h.dataService.UpdateMediaBooleans(ctx, mb)
	if err != nil {
		return nil, err
	}

	return res, nil
}
