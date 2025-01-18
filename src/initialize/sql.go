package initialize

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"mockserver.jiratviriyataranon.io/src/config"
)

func SqlPool(ctx context.Context, getEnv func(string) string) (*sql.DB, error) {
	sqlConfig, err := config.Sql(getEnv)
	if err != nil {
		return nil, fmt.Errorf("Error getting SQL configs: %w", err)
	}

	fmt.Printf("Connecting to SQL at %v\n", sqlConfig.Address)
	sqlPool, err := sql.Open("mysql", sqlConfig.Address)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to SQL: %w", err)
	}
	context.AfterFunc(ctx, func() { sqlPool.Close() }) // On close error, the resource will eventually be deallocated by the kernel

	sqlPool.SetMaxIdleConns(sqlConfig.MaxIdleConns)
	sqlPool.SetMaxOpenConns(sqlConfig.MaxOpenConns)
	sqlPool.SetConnMaxIdleTime(sqlConfig.MaxIdleTime)
	sqlPool.SetConnMaxLifetime(sqlConfig.MaxLifeTime)

	return sqlPool, nil
}
