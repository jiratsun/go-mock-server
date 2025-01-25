package config

import (
	"fmt"
	"strconv"
	"time"
)

func Sql(getEnv func(string) string) (sqlConfig, error) {
	address := getEnv("SQL_SOURCE_NAME")

	maxIdleConns, err := strconv.Atoi(getEnv("SQL_MAX_IDLE_CONNS"))
	if err != nil {
		return sqlConfig{}, fmt.Errorf("Error parsing SQL_MAX_IDLE_CONNS: %w", err)
	}

	maxOpenConns, err := strconv.Atoi(getEnv("SQL_MAX_OPEN_CONNS"))
	if err != nil {
		return sqlConfig{}, fmt.Errorf("Error parsing SQL_MAX_OPEN_CONNS: %w", err)
	}

	maxIdleTime, err := time.ParseDuration(getEnv("SQL_MAX_IDLE_TIME"))
	if err != nil {
		return sqlConfig{}, fmt.Errorf("Error parsing SQL_MAX_IDLE_TIME: %w", err)
	}

	maxLifeTime, err := time.ParseDuration(getEnv("SQL_MAX_LIFE_TIME"))
	if err != nil {
		return sqlConfig{}, fmt.Errorf("Error parsing SQL_MAX_LIFE_TIME: %w", err)
	}

	initialConnectTimeout, err := time.ParseDuration(getEnv("SQL_INITIAL_CONNECT_TIMEOUT"))
	if err != nil {
		return sqlConfig{}, fmt.Errorf("Error parsing SQL_INITIAL_CONNECT_TIMEOUT: %w", err)
	}

	return sqlConfig{
		Address:               address,
		MaxIdleConns:          maxIdleConns,
		MaxOpenConns:          maxOpenConns,
		MaxIdleTime:           maxIdleTime,
		MaxLifeTime:           maxLifeTime,
		InitialConnectTimeout: initialConnectTimeout,
	}, nil
}
