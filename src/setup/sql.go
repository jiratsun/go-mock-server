package setup

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"mockserver.jiratviriyataranon.io/src/time"
)

func SqlPool(ctx context.Context, getEnv func(string) string) (*sql.DB, error) {
	address := getEnv("SQL_SOURCE_NAME")

	fmt.Printf("Connecting to SQL at %v\n", address)
	sqlPool, err := sql.Open("mysql", address)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to SQL: %w", err)
	}
	context.AfterFunc(ctx, func() { sqlPool.Close() }) // On close error, the resource will eventually be deallocated by the kernel

	sqlConfig, err := getSqlConfig(getEnv)
	if err != nil {
		return nil, fmt.Errorf("Error getting SQL configs: %w", err)
	}

	sqlPool.SetMaxIdleConns(sqlConfig.maxIdleConns)
	sqlPool.SetMaxOpenConns(sqlConfig.maxOpenConns)
	sqlPool.SetConnMaxIdleTime(time.OfSeconds(sqlConfig.maxIdleTimeSeconds))
	sqlPool.SetConnMaxLifetime(time.OfSeconds(sqlConfig.maxLifeTimeSeconds))

	return sqlPool, nil
}

type sqlConfig struct {
	maxIdleConns       int
	maxOpenConns       int
	maxIdleTimeSeconds int
	maxLifeTimeSeconds int
}

func getSqlConfig(getEnv func(string) string) (sqlConfig, error) {
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

	return sqlConfig{maxIdleConns, maxOpenConns, maxIdleTimeSeconds, maxLifeTimeSeconds}, nil
}
