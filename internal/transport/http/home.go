package http

import (
	"net/http"

	"github.com/rs/zerolog"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	data := map[string]string{}

	err := h.template.ExecuteTemplate(w, "index.html", data)

	if err != nil {
		log.Error().Err(err).Msg("template execute failed")
	}
}
