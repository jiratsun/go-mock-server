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

	"github.com/joho/godotenv"
	timeutil "mockserver.jiratviriyataranon.io/src/util"
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
	address := net.JoinHostPort(getEnv("SERVER_HOST"), getEnv("SERVER_PORT"))
	serverCtx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	server := &http.Server{
		Addr:        address,
		BaseContext: func(net.Listener) context.Context { return serverCtx },
	}

	gracePeriod, err := strconv.Atoi(getEnv("SHUTDOWN_GRACE_PERIOD_SECONDS"))
	if err != nil {
		return fmt.Errorf("Error parsing environment variables: %w", err)
	}

	shutdownErr := make(chan error)
	go func() {
		<-serverCtx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), timeutil.OfSeconds(gracePeriod))
		defer cancel()

		fmt.Println("Shutting down server")
		shutdownErr <- server.Shutdown(shutdownCtx)
	}()

	fmt.Printf("Starting server at %v\n", address)
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
