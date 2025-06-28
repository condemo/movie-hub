package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MovieHandler struct{}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler{}
}

func (h *MovieHandler) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", h.GetMovies)
	return r
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Movies")
}
