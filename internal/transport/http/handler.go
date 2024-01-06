package http

import (
	"net/http"
)

type Handler struct {
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
