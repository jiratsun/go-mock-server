package initializer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"mockserver.jiratviriyataranon.io/src/core/config"
)

func InitHandler(ctx context.Context, getEnv func(string) string) (http.Handler, error) {
	sqlPool, err := initSqlPool(ctx, getEnv)
	if err != nil {
		return nil, fmt.Errorf("Error setting up SQL: %w", err)
	}

	hostStore := &config.HostStore{SqlPool: sqlPool, GetEnv: getEnv}
	pathStore := &config.PathStore{SqlPool: sqlPool, GetEnv: getEnv}

	hostHandler := &config.HostHandler{Store: hostStore}
	pathHandler := &config.PathHandler{Store: pathStore}

	return initRoute(
		chi.NewRouter(),
		hostHandler,
		pathHandler,
	), nil
}
