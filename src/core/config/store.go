package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"mockserver.jiratviriyataranon.io/.jet-gen/go_mock_server/model"
	. "mockserver.jiratviriyataranon.io/.jet-gen/go_mock_server/table"
)

type ConfigStore struct {
	SqlPool *sql.DB
	GetEnv  func(string) string
}

func (store *ConfigStore) findAllHost(ctx context.Context) ([]model.Host, error) {
	result := make([]model.Host, 0)
	statement := Host.SELECT(Host.AllColumns)

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

func (store *ConfigStore) upsertManyHost(ctx context.Context, hosts []hostUpsertMany) error {
	statement := Host.
		INSERT(Host.DomainName, Host.Alias_, Host.Description).
		MODELS(hosts)

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

func (store *ConfigStore) upsertManyPath(ctx context.Context, paths []pathUpsertMany) error {
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
