package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"mockserver.jiratviriyataranon.io/src/config"
	"mockserver.jiratviriyataranon.io/src/core/path"
	"mockserver.jiratviriyataranon.io/src/initialize"
)

func main() {
	err := godotenv.Load(fmt.Sprintf(".%v.env", os.Args[1]))
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		os.Exit(1)
	}

	err = run(context.Background(), os.Getenv)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run(ctx context.Context, getEnv func(string) string) error {
	serverConfig, err := config.Server(getEnv)
	if err != nil {
		return fmt.Errorf("Error getting server configs: %w", err)
	}

	serverCtx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	handler, err := initializeHandler(serverCtx, getEnv)
	if err != nil {
		return fmt.Errorf("Error setting up handler: %w", err)
	}

	server := &http.Server{
		Addr:        serverConfig.Address,
		BaseContext: func(net.Listener) context.Context { return serverCtx },
		Handler:     handler,
	}

	shutdownErr := make(chan error)

	go func() {
		<-serverCtx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), serverConfig.GracePeriod)
		defer cancel()

		fmt.Println("Shutting down server")
		shutdownErr <- server.Shutdown(shutdownCtx)
	}()

	fmt.Printf("Starting server at %v\n", serverConfig.Address)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("Error starting server: %w", err)
	}

	err = <-shutdownErr
	if err != nil {
		return fmt.Errorf("Error shutting down server: %w", err)
	}

	fmt.Println("Server shut down")
	return nil
}

func initializeHandler(ctx context.Context, getEnv func(string) string) (http.Handler, error) {
	sqlPool, err := initialize.SqlPool(ctx, getEnv)
	if err != nil {
		return nil, fmt.Errorf("Error setting up SQL: %w", err)
	}

	pathStore := &path.PathStore{SqlPool: sqlPool}

	pathHandler := &path.PathHandler{Store: pathStore}

	return route(
		chi.NewRouter(),
		pathHandler,
	), nil
}
