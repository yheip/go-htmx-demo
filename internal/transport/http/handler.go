package http

import (
	"embed"
	"html/template"
)

//go:embed templates/*
var htmlFiles embed.FS

type Handler struct {
	template *template.Template
}

func NewHandler() *Handler {
	return &Handler{
		template: template.Must(template.ParseFS(htmlFiles, "templates/*")),
	}
}
