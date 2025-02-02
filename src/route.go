package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"mockserver.jiratviriyataranon.io/src/core/host"
	"mockserver.jiratviriyataranon.io/src/core/path"
)

func route(
	router chi.Router,
	hostHandler *host.HostHandler,
	pathHandler *path.PathHandler,
) http.Handler {
	hostRouter := chi.NewRouter()
	hostRouter.Get("/", hostHandler.HandleGet)
	hostRouter.Post("/", hostHandler.HandleRegisterHost)

	pathRouter := chi.NewRouter()
	pathRouter.Delete("/", pathHandler.HandleDelete)
	pathRouter.Get("/", pathHandler.HandleGet)
	pathRouter.Post("/", pathHandler.HandleRegisterPathToHost)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/config/host", hostRouter)
		r.Mount("/config/path", pathRouter)
	})
	return router
}
