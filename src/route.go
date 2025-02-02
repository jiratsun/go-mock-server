package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"mockserver.jiratviriyataranon.io/src/core/host"
)

func route(
	router chi.Router,
	hostHandler *host.HostHandler,
) http.Handler {
	hostRouter := chi.NewRouter().Group(func(r chi.Router) {
		r.Get("/", hostHandler.HandleGet)
		r.Post("/", hostHandler.HandleRegisterHost)
	})

	pathRouter := chi.NewRouter().Group(func(r chi.Router) {
		r.Post("/", hostHandler.HandleRegisterPathToHost)
	})

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/config/host", hostRouter)
		r.Mount("/config/path", pathRouter)
	})
	return router
}
