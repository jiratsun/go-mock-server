package config

import (
	"fmt"
	"net"
	"strconv"

	"mockserver.jiratviriyataranon.io/src/time"
)

func Server(getEnv func(string) string) (serverConfig, error) {
	address := net.JoinHostPort(getEnv("SERVER_HOST"), getEnv("SERVER_PORT"))

	gracePeriod, err := strconv.Atoi(getEnv("SHUTDOWN_GRACE_PERIOD_SECONDS"))
	if err != nil {
		return serverConfig{}, fmt.Errorf("Error parsing environment variables: %w", err)
	}

	return serverConfig{
		Address:     address,
		GracePeriod: time.OfSeconds(gracePeriod),
	}, nil
}
