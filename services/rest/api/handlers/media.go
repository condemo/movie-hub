package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/condemo/movie-hub/services/common/config"
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
	r.Get("/", MakeHandler(h.getLastUpdates))
	r.Get("/{id}", MakeHandler(h.getOneMedia))
	r.Put("/", MakeHandler(h.updateMedia))
	r.Put("/resume", MakeHandler(h.updateByMediaResume))
	r.Delete("/{id}", MakeHandler(h.deleteMedia))
	return r
}

func (h *MediaHandler) getLastUpdates(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	limit := new(int32)
	l, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	if err == nil {
		*limit = int32(l)
	} else {
		*limit = config.General.DefaultDataLimit
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

func (h *MediaHandler) getOneMedia(w http.ResponseWriter, r *http.Request) error {
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

func (h *MediaHandler) updateMedia(w http.ResponseWriter, r *http.Request) error {
	media := new(pb.Media)
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*8)
	defer cancel()

	err := json.NewDecoder(r.Body).Decode(media)
	if err != nil {
		return err
	}

	res, err := h.dataConn.UpdateMedia(ctx, &pb.UpdateMediaReq{
		Media: media,
	})
	if err != nil {
		return err
	}

	JsonResponse(w, http.StatusOK, res.GetMedia())

	return nil
}

func (h *MediaHandler) deleteMedia(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	_, err := h.dataConn.DeleteMedia(ctx, &pb.MediaRequest{})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}

func (h *MediaHandler) updateByMediaResume(w http.ResponseWriter, r *http.Request) error {
	mc := pb.MediaUpdateBool{}
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	err := json.NewDecoder(r.Body).Decode(&mc)
	if err != nil {
		return err
	}

	res, err := h.dataConn.UpdateMediaBooleans(ctx, &mc)
	if err != nil {
		return err
	}

	JsonResponse(w, http.StatusOK, res)

	return nil
}
