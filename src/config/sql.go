package config

import (
	"fmt"
	"strconv"
)

func Sql(getEnv func(string) string) (sqlConfig, error) {
	address := getEnv("SQL_SOURCE_NAME")

	maxIdleConns, err := strconv.Atoi(getEnv("SQL_MAX_IDLE_CONNS"))
	if err != nil {
		return sqlConfig{}, fmt.Errorf("Error parsing environment variables: %w", err)
	}

	maxOpenConns, err := strconv.Atoi(getEnv("SQL_MAX_OPEN_CONNS"))
	if err != nil {
		return sqlConfig{}, fmt.Errorf("Error parsing environment variables: %w", err)
	}

	maxIdleTimeSeconds, err := strconv.Atoi(getEnv("SQL_MAX_IDLE_TIME_SECONDS"))
	if err != nil {
		return sqlConfig{}, fmt.Errorf("Error parsing environment variables: %w", err)
	}

	maxLifeTimeSeconds, err := strconv.Atoi(getEnv("SQL_MAX_LIFE_TIME_SECONDS"))
	if err != nil {
		return sqlConfig{}, fmt.Errorf("Error parsing environment variables: %w", err)
	}

	return sqlConfig{
		Address:            address,
		MaxIdleConns:       maxIdleConns,
		MaxOpenConns:       maxOpenConns,
		MaxIdleTimeSeconds: maxIdleTimeSeconds,
		MaxLifeTimeSeconds: maxLifeTimeSeconds,
	}, nil
}
