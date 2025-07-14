package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	r.Get("/", h.GetLastUpdates)
	return r
}

func (h *MediaHandler) GetLastUpdates(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var limit int32
	l, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	if err == nil {
		limit = int32(l)
	}

	data, err := h.dataConn.GetLastUpdates(ctx, &pb.LastUpdatesRequest{
		Type:  pb.MediaType_Both,
		Limit: limit,
	})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.GetMediaList())
}

func (h *MediaHandler) GetOneMedia(w http.ResponseWriter, r *http.Request) {
	// TODO:
}
