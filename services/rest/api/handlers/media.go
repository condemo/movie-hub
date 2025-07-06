package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

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
	r.Get("/", h.GetMovies)
	return r
}

func (h *MediaHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	data, err := h.dataConn.GetLastUpdates(ctx, &pb.LastUpdatesRequest{Type: pb.MediaType_Movie})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, data.GetMsg())
}
