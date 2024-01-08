package http

import "github.com/go-chi/chi/v5"

func AddRoutes(router chi.Router) {

	handler := NewHandler()

	router.Get("/", handler.Home)
	router.Get("/hello", handler.Hello)
	router.Get("/sse", handler.HandleSSE)
	router.Get("/poll", handler.Poll)
	router.Get("/update", handler.GetUpdate)
	router.Get("/redirect", handler.GetRedirect)
	router.Get("/bye", handler.Bye)
	router.Post("/code", handler.NewCode)
}
