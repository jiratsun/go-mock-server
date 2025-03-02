package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"mockserver.jiratviriyataranon.io/src/core/config"
)

func route(
	router chi.Router,
	hostHandler *config.HostHandler,
	pathHandler *config.PathHandler,
) http.Handler {
	hostRouter := chi.NewRouter().Group(func(r chi.Router) {
		r.Get("/", hostHandler.HandleGetHost)
		r.Post("/", hostHandler.HandleRegisterHost)
		r.Delete("/", hostHandler.HandleDeleteHost)
	})

	pathRouter := chi.NewRouter().Group(func(r chi.Router) {
		r.Get("/", pathHandler.HandleGetPath)
		r.Post("/", pathHandler.HandleRegisterPath)
	})

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/config/host", hostRouter)
		r.Mount("/config/path", pathRouter)
	})
	return router
}
