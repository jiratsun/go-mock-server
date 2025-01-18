package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	timeutil "mockserver.jiratviriyataranon.io/src/util"
)

func main() {
	err := run(context.Background(), os.Getenv)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(ctx context.Context, getEnv func(string) string) error {
	serverCtx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	server := &http.Server{
		Addr:        net.JoinHostPort(getEnv("SERVER_HOST"), getEnv("SERVER_PORT")),
		BaseContext: func(net.Listener) context.Context { return serverCtx },
	}

	shutdownErr := make(chan error)

	gracePeriod, err := strconv.Atoi(getEnv("SHUTDOWN_GRACE_PERIOD"))
	if err != nil {
		return fmt.Errorf("Error while parsing environment variables: %w", err)
	}

	go func() {
		<-serverCtx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), timeutil.OfSeconds(gracePeriod))
		defer cancel()

		fmt.Println("Shutting down server")
		shutdownErr <- server.Shutdown(shutdownCtx)
	}()

	fmt.Printf("Starting server at %v\n", net.JoinHostPort(getEnv("SERVER_HOST"), getEnv("SERVER_PORT")))
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("Error while starting server: %w", err)
	}

	err = <-shutdownErr
	if err != nil {
		return fmt.Errorf("Error while shutting down server: %w", err)
	}

	fmt.Println("Server shut down")
	return nil
}
