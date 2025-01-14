package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	timeutil "mockserver.jiratviriyataranon.io/src/util"
)

func main() {
	serverCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	server := &http.Server{}

	shutdownErr := make(chan error)

	go func() {
		<-serverCtx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), timeutil.OfSeconds(5))
		defer cancel()

		fmt.Println("Shutting down server")
		shutdownErr <- server.Shutdown(shutdownCtx)
	}()

	fmt.Println("Starting server")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Error while starting server")
	}

	err = <-shutdownErr
	if err != nil {
		fmt.Println("Error while shutting down server")
	}

	fmt.Println("Server shut down")
}
