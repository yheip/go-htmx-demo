package http

import "github.com/go-chi/chi/v5"

func AddRoutes(router chi.Router) {

	handler := NewHandler()

	router.Get("/", handler.Home)
	router.Get("/hello", handler.Hello)
	router.Get("/sse", handler.HandleSSE)
	router.Post("/code", handler.NewCode)
}
