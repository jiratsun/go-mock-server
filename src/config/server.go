package config

import (
	"fmt"
	"net"
	"time"
)

func Server(getEnv func(string) string) (serverConfig, error) {
	address := net.JoinHostPort(getEnv("SERVER_HOST"), getEnv("SERVER_PORT"))

	gracePeriod, err := time.ParseDuration(getEnv("SHUTDOWN_GRACE_PERIOD"))
	if err != nil {
		return serverConfig{}, fmt.Errorf("Error parsing SHUTDOWN_GRACE_PERIOD: %w", err)
	}

	return serverConfig{
		Address:     address,
		GracePeriod: gracePeriod,
	}, nil
}
