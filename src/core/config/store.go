package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	i "mockserver.jiratviriyataranon.io/src/integers"
	"mockserver.jiratviriyataranon.io/src/strings"
)

type ConfigStore struct {
	SqlPool *sql.DB
	GetEnv  func(string) string
}

func (store *ConfigStore) findAllWithPath(ctx context.Context) ([]hostWithPath, error) {
	var query strings.StringBuilder
	query.WriteStringln("SELECT * FROM host")
	query.WriteString("JOIN path_to_host ON host.alias = path_to_host.host_alias")
	result := make([]hostWithPath, 0)

	timeout, err := time.ParseDuration(store.GetEnv("SQL_READ_TIMEOUT"))
	if err != nil {
		return result, fmt.Errorf("Error parsing SQL_READ_TIMEOUT: %w", err)
	}

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	rows, err := store.SqlPool.QueryContext(queryCtx, query.String())
	if err != nil {
		return result, fmt.Errorf("Error reading from database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var hostWithPath hostWithPath

		err := rows.Scan(
			&hostWithPath.id,
			&hostWithPath.host,
			&hostWithPath.alias,
			&hostWithPath.isActive,
			&hostWithPath.createdAt,
			&hostWithPath.updatedAt,
			&hostWithPath.path_id,
			&hostWithPath.path_path,
			&hostWithPath.path_hostAlias,
			&hostWithPath.path_isActive,
			&hostWithPath.path_createdAt,
			&hostWithPath.path_updatedAt,
		)
		if err != nil {
			return result, fmt.Errorf("Error parsing rows: %w", err)
		}

		result = append(result, hostWithPath)
	}

	return result, nil
}

func (store *ConfigStore) upsertMany(ctx context.Context, aliasToHost []aliasToHostUpsertMany) error {
	var query strings.StringBuilder
	query.WriteStringln("INSERT INTO host (alias, host) VALUES")
	query.WriteStringlnRepeat("(?, ?),", i.Dec(len(aliasToHost)))
	query.WriteStringln("(?, ?) AS new")
	query.WriteString("ON DUPLICATE KEY UPDATE host=new.host;")

	var args []any
	for _, row := range aliasToHost {
		args = append(args, row.alias, row.host)
	}

	timeout, err := time.ParseDuration(store.GetEnv("SQL_WRITE_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("Error parsing SQL_WRITE_TIMEOUT: %w", err)
	}

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err = store.SqlPool.QueryRowContext(queryCtx, query.String(), args...).Err()
	if err != nil {
		return fmt.Errorf("Error writing to database: %w", err)
	}

	return nil
}

func (store *ConfigStore) upsertManyPath(ctx context.Context, pathToHost []pathToHostUpsertMany) error {
	var query strings.StringBuilder
	query.WriteStringln("INSERT INTO path_to_host (path, host_alias) VALUES")
	query.WriteStringlnRepeat("(?, ?),", i.Dec(len(pathToHost)))
	query.WriteStringln("(?, ?) AS new")
	query.WriteString("ON DUPLICATE KEY UPDATE host_alias=new.host_alias;")

	var args []any
	for _, row := range pathToHost {
		args = append(args, row.path, row.hostAlias)
	}

	timeout, err := time.ParseDuration(store.GetEnv("SQL_WRITE_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("Error parsing SQL_WRITE_TIMEOUT: %w", err)
	}

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err = store.SqlPool.QueryRowContext(queryCtx, query.String(), args...).Err()
	if err != nil {
		return fmt.Errorf("Error writing to database: %w", err)
	}

	return nil
}
