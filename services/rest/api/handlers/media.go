package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/condemo/movie-hub/services/common/errs"
	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type MediaHandler struct {
	dataConn pb.DataHandlerClient
}

func NewMediaHandler(dataConn *grpc.ClientConn) *MediaHandler {
	dc := pb.NewDataHandlerClient(dataConn)
	return &MediaHandler{
		dataConn: dc,
	}
}

func (h *MediaHandler) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", MakeHandler(h.GetLastUpdates))
	r.Get("/{id}", MakeHandler(h.GetOneMedia))
	return r
}

func (h *MediaHandler) GetLastUpdates(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	limit := new(int32)
	l, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	if err == nil {
		*limit = int32(l)
	}

	filter := r.URL.Query().Get("filter")
	if filter != "" {
		media, err := h.dataConn.GetMediaFiltered(ctx, &pb.MediaFilteredRequest{
			Filter: pb.FilterBy(pb.FilterBy_value[filter]),
			Limit:  limit,
		})
		if err != nil {
			return err
		}
		if media.GetMediaList() == nil {
			JsonResponse(w, http.StatusNotFound, "media not found")
			return nil
		}
		JsonResponse(w, http.StatusOK, media.GetMediaList())
		return nil
	}

	data, err := h.dataConn.GetLastUpdates(ctx, &pb.LastUpdatesRequest{
		Type:  pb.MediaType_Both,
		Limit: limit,
	})
	if err != nil {
		return err
	}

	JsonResponse(w, http.StatusOK, data.GetMediaList())

	return nil
}

func (h *MediaHandler) GetOneMedia(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return errs.NewApiError(err, http.StatusBadRequest, "invalid media id")
	}

	res, err := h.dataConn.GetOneMedia(ctx, &pb.MediaRequest{Id: id})
	if err != nil {
		return err
	}

	JsonResponse(w, http.StatusOK, res)

	return nil
}
