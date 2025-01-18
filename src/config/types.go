package config

import "time"

type serverConfig struct {
	Address     string
	GracePeriod time.Duration
}

type sqlConfig struct {
	MaxIdleConns       int
	MaxOpenConns       int
	MaxIdleTimeSeconds int
	MaxLifeTimeSeconds int
}
