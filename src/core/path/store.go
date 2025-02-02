package path

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	i "mockserver.jiratviriyataranon.io/src/integers"
	"mockserver.jiratviriyataranon.io/src/strings"
)

type PathStore struct {
	SqlPool *sql.DB
	GetEnv  func(string) string
}

func (store *PathStore) upsertMany(ctx context.Context, pathToHost []pathToHostUpsertMany) error {
	var query strings.StringBuilder
	query.WriteStringln("INSERT INTO path_to_host (path, host) VALUES")
	query.WriteStringlnRepeat("(?, ?),", i.Dec(len(pathToHost)))
	query.WriteStringln("(?, ?) AS new")
	query.WriteString("ON DUPLICATE KEY UPDATE host=new.host;")

	var args []any
	for _, row := range pathToHost {
		args = append(args, row.path, row.host)
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

func (store *PathStore) findAll(ctx context.Context) ([]pathToHost, error) {
	query := "SELECT * FROM path_to_host"
	result := make([]pathToHost, 0)

	timeout, err := time.ParseDuration(store.GetEnv("SQL_READ_TIMEOUT"))
	if err != nil {
		return result, fmt.Errorf("Error parsing SQL_READ_TIMEOUT: %w", err)
	}

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	rows, err := store.SqlPool.QueryContext(queryCtx, query)
	if err != nil {
		return result, fmt.Errorf("Error reading from database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var pathToHost pathToHost

		err := rows.Scan(
			&pathToHost.id,
			&pathToHost.path,
			&pathToHost.host_alias,
			&pathToHost.isActive,
			&pathToHost.createdAt,
			&pathToHost.updatedAt,
		)
		if err != nil {
			return result, fmt.Errorf("Error parsing rows: %w", err)
		}

		result = append(result, pathToHost)
	}

	return result, nil
}
