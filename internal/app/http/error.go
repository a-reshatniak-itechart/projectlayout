package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/a-reshatniak-itechart/projectlayout/internal/app"
)

func handleErr(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, app.ErrUserNotFound):
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
