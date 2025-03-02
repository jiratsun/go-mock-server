package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/mysql"
	m "mockserver.jiratviriyataranon.io/.jet-gen/go_mock_server/model"
	. "mockserver.jiratviriyataranon.io/.jet-gen/go_mock_server/table"
	"mockserver.jiratviriyataranon.io/src/data"
)

type PathStore struct {
	SqlPool *sql.DB
	GetEnv  func(string) string
}

func (store *PathStore) findAllPath(ctx context.Context) ([]m.Path, error) {
	result := make([]m.Path, 0)
	statement := Path.SELECT(Path.AllColumns)

	timeout, err := time.ParseDuration(store.GetEnv("SQL_READ_TIMEOUT"))
	if err != nil {
		return result, fmt.Errorf("Error parsing SQL_READ_TIMEOUT: %w", err)
	}

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err = statement.QueryContext(queryCtx, store.SqlPool, &result)
	if err != nil {
		return result, fmt.Errorf("Error reading from database: %w", err)
	}

	return result, nil
}

func (store *PathStore) upsertManyPath(ctx context.Context, paths []pathUpsertMany) error {
	statement := Path.
		INSERT(Path.Path, Path.DefaultHost, Path.Description).
		MODELS(paths).
		AS_NEW().
		ON_DUPLICATE_KEY_UPDATE(
			Path.DefaultHost.SET(Path.NEW.DefaultHost),
			Path.Description.SET(Path.NEW.Description),
		)

	timeout, err := time.ParseDuration(store.GetEnv("SQL_WRITE_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("Error parsing SQL_WRITE_TIMEOUT: %w", err)
	}

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err = statement.ExecContext(queryCtx, store.SqlPool)
	if err != nil {
		return fmt.Errorf("Error writing to database: %w", err)
	}

	return nil
}

func (store *PathStore) deleteManyPath(ctx context.Context, paths pathModifyMany) error {
	whereClause := mysql.Bool(false)

	if len(paths.Path) > 0 {
		domainName := data.Map(paths.Path, func(v string) mysql.Expression {
			return mysql.String(v)
		})
		whereClause = whereClause.OR(Path.Path.IN(domainName...))
	}

	statement := Path.DELETE().WHERE(whereClause)

	timeout, err := time.ParseDuration(store.GetEnv("SQL_WRITE_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("Error parsing SQL_WRITE_TIMEOUT: %w", err)
	}

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err = statement.ExecContext(queryCtx, store.SqlPool)
	if err != nil {
		return fmt.Errorf("Error writing to database: %w", err)
	}

	return nil
}

func (store *PathStore) toggleManyPath(ctx context.Context, paths pathModifyMany, enable bool) error {
	update := m.Path{IsActive: enable}
	whereClause := mysql.Bool(false)

	if len(paths.Path) > 0 {
		domainName := data.Map(paths.Path, func(v string) mysql.Expression {
			return mysql.String(v)
		})
		whereClause = whereClause.OR(Path.Path.IN(domainName...))
	}

	statement := Path.UPDATE(Path.IsActive).MODEL(update).WHERE(whereClause)

	timeout, err := time.ParseDuration(store.GetEnv("SQL_WRITE_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("Error parsing SQL_WRITE_TIMEOUT: %w", err)
	}

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err = statement.ExecContext(queryCtx, store.SqlPool)
	if err != nil {
		return fmt.Errorf("Error writing to database: %w", err)
	}

	return nil
}
