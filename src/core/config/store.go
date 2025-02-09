package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/mysql"
	. "mockserver.jiratviriyataranon.io/.jet-gen/go_mock_server/table"
)

type ConfigStore struct {
	SqlPool *sql.DB
	GetEnv  func(string) string
}

func (store *ConfigStore) findAllWithPath(ctx context.Context) ([]hostWithPath, error) {
	result := make([]hostWithPath, 0)
	statement := mysql.
		SELECT(Host.AllColumns, PathToHost.AllColumns).
		FROM(Host.
			LEFT_JOIN(PathToHost, Host.Alias_.EQ(PathToHost.HostAlias)),
		)

	timeout, err := time.ParseDuration(store.GetEnv("SQL_READ_TIMEOUT"))
	if err != nil {
		return result, fmt.Errorf("Error parsing SQL_READ_TIMEOUT: %w", err)
	}

	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	rows, err := statement.Rows(queryCtx, store.SqlPool)
	if err != nil {
		return result, fmt.Errorf("Error reading from database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var hostWithPath hostWithPath

		err := rows.Rows.Scan(
			&hostWithPath.id,
			&hostWithPath.host,
			&hostWithPath.alias,
			&hostWithPath.isActive,
			&hostWithPath.createdAt,
			&hostWithPath.updatedAt,
			&hostWithPath.pathId,
			&hostWithPath.path,
			&hostWithPath.hostAlias,
			&hostWithPath.pathIsActive,
			&hostWithPath.pathCreatedAt,
			&hostWithPath.pathUpdatedAt,
		)
		if err != nil {
			return result, fmt.Errorf("Error parsing rows: %w", err)
		}

		result = append(result, hostWithPath)
	}

	return result, nil
}

func (store *ConfigStore) upsertMany(ctx context.Context, aliasToHost []aliasToHostUpsertMany) error {
	statement := Host.
		INSERT(Host.Alias_, Host.Host).
		MODELS(aliasToHost).
		AS_NEW().
		ON_DUPLICATE_KEY_UPDATE(Host.Host.SET(Host.NEW.Host))

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

func (store *ConfigStore) upsertManyPath(ctx context.Context, pathToHost []pathToHostUpsertMany) error {
	statement := PathToHost.
		INSERT(PathToHost.Path, PathToHost.HostAlias).
		MODELS(pathToHost).
		AS_NEW().
		ON_DUPLICATE_KEY_UPDATE(PathToHost.HostAlias.SET(PathToHost.NEW.HostAlias))

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
