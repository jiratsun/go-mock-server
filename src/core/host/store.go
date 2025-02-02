package host

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	i "mockserver.jiratviriyataranon.io/src/integers"
	"mockserver.jiratviriyataranon.io/src/strings"
)

type HostStore struct {
	SqlPool *sql.DB
	GetEnv  func(string) string
}

func (store *HostStore) findAll(ctx context.Context) ([]host, error) {
	query := "SELECT * FROM host"
	result := make([]host, 0)

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
		var aliasToHost host

		err := rows.Scan(
			&aliasToHost.id,
			&aliasToHost.host,
			&aliasToHost.alias,
			&aliasToHost.isActive,
			&aliasToHost.createdAt,
			&aliasToHost.updatedAt,
		)
		if err != nil {
			return result, fmt.Errorf("Error parsing rows: %w", err)
		}

		result = append(result, aliasToHost)
	}

	return result, nil
}

func (store *HostStore) upsertMany(ctx context.Context, aliasToHost []aliasToHostUpsertMany) error {
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
