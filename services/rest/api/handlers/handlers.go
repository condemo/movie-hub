package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type customHandler func(http.ResponseWriter, *http.Request) error

func MakeHandler(f customHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// TODO: crear un error filter usando un switch err := err.(type)

			// FIX: Borrar, crear sistema centralizado de errores
			fmt.Println(err)
			JsonResponse(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func JsonResponse(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
