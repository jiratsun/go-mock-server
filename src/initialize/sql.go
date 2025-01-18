package initialize

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"mockserver.jiratviriyataranon.io/src/config"
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

	sqlConfig, err := config.Sql(getEnv)
	if err != nil {
		return nil, fmt.Errorf("Error getting SQL configs: %w", err)
	}

	sqlPool.SetMaxIdleConns(sqlConfig.MaxIdleConns)
	sqlPool.SetMaxOpenConns(sqlConfig.MaxOpenConns)
	sqlPool.SetConnMaxIdleTime(time.OfSeconds(sqlConfig.MaxIdleTimeSeconds))
	sqlPool.SetConnMaxLifetime(time.OfSeconds(sqlConfig.MaxLifeTimeSeconds))

	return sqlPool, nil
}
