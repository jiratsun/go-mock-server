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
