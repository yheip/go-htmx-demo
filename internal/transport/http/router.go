package http

import "github.com/go-chi/chi/v5"

func AddRoutes(router chi.Router) {
	handler := Handler{}

	router.Get("/", handler.Hello)
}
