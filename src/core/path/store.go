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
