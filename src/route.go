package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"mockserver.jiratviriyataranon.io/src/core/path"
)

func route(
	router chi.Router,
	pathHandler *path.PathHandler,
) http.Handler {
	configRouter := chi.NewRouter()
	configRouter.Route("/path", func(r chi.Router) {
		r.Delete("/", pathHandler.HandleDelete)
		r.Get("/", pathHandler.HandleGet)
		r.Post("/", pathHandler.HandleRegisterPathToHost)
	})

	router.Mount("/v1/config", configRouter)
	return router
}
