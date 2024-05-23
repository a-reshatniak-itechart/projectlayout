package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/a-reshatniak-itechart/projectlayout/internal/app"
)

func NewUserHandler(c app.UserController) UserHandler {
	return UserHandler{c: c}
}

type UserHandler struct {
	c app.UserController
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.c.GetUser(r.Context(), id)
	if err != nil {
		handleErr(w, err)
		return
	}

	data, err := json.Marshal(u)
	if err != nil {
		handleErr(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
