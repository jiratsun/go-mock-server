package initializer

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"mockserver.jiratviriyataranon.io/src/core/config"
)

func initRoute(
	router chi.Router,
	hostHandler *config.HostHandler,
	pathHandler *config.PathHandler,
) http.Handler {
	hostRouter := chi.NewRouter().Group(func(r chi.Router) {
		r.Get("/", hostHandler.HandleGetHost)
		r.Post("/", hostHandler.HandleRegisterHost)
		r.Delete("/", hostHandler.HandleDeleteHost)
		r.Get("/enable", hostHandler.HandleGetActiveHost)
		r.Post("/enable", hostHandler.HandleEnableHost)
		r.Get("/disable", hostHandler.HandleGetInactiveHost)
		r.Post("/disable", hostHandler.HandleDisableHost)
	})

	pathRouter := chi.NewRouter().Group(func(r chi.Router) {
		r.Get("/", pathHandler.HandleGetPath)
		r.Post("/", pathHandler.HandleRegisterPath)
		r.Delete("/", pathHandler.HandleDeletePath)
		r.Get("/enable", pathHandler.HandleGetActivePath)
		r.Post("/enable", pathHandler.HandleEnablePath)
		r.Get("/disable", pathHandler.HandleGetInactivePath)
		r.Post("/disable", pathHandler.HandleDisablePath)
	})

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/config/host", hostRouter)
		r.Mount("/config/path", pathRouter)
	})
	return router
}
