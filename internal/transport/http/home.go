package http

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/rs/zerolog"
)

//go:embed templates/*
var htmlFiles embed.FS

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())

	t, err := template.ParseFS(htmlFiles, "templates/index.html")
	if err != nil {
		log.Error().Err(err).Msg("template parse failed")
	}

	data := map[string]string{}

	err = t.Execute(w, data)

	if err != nil {
		log.Error().Err(err).Msg("template execute failed")
	}
}
