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

type HostStore struct {
	SqlPool *sql.DB
	GetEnv  func(string) string
}

func (store *HostStore) findAllHost(ctx context.Context, isActive *bool) ([]m.Host, error) {
	result := make([]m.Host, 0)
	statement := Host.SELECT(Host.AllColumns)

	if isActive != nil {
		statement = statement.WHERE(Host.IsActive.EQ(mysql.Bool(*isActive)))
	}

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

func (store *HostStore) upsertManyHost(ctx context.Context, hosts []hostUpsertMany) error {
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

func (store *HostStore) deleteManyHost(ctx context.Context, hosts hostModifyMany) error {
	whereClause := mysql.Bool(false)

	if len(hosts.DomainName) > 0 {
		domainName := data.Map(hosts.DomainName, func(v string) mysql.Expression {
			return mysql.String(v)
		})
		whereClause = whereClause.OR(Host.DomainName.IN(domainName...))
	}

	if len(hosts.Alias) > 0 {
		alias := data.Map(hosts.Alias, func(v string) mysql.Expression {
			return mysql.String(v)
		})
		whereClause = whereClause.OR(Host.Alias_.IN(alias...))
	}

	if len(hosts.Both) > 0 {
		data.ForEach(hosts.Both, func(both data.Tuple2[string, string]) {
			criteria1 := Host.DomainName.EQ(mysql.String(both.Left))
			criteria2 := Host.Alias_.EQ(mysql.String(both.Right))
			whereClause = whereClause.OR(criteria1.AND(criteria2))
		})
	}

	statement := Host.DELETE().WHERE(whereClause)

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

func (store *HostStore) toggleManyHost(ctx context.Context, hosts hostModifyMany, enable bool) error {
	update := m.Host{IsActive: enable}
	whereClause := mysql.Bool(false)

	if len(hosts.DomainName) > 0 {
		domainName := data.Map(hosts.DomainName, func(v string) mysql.Expression {
			return mysql.String(v)
		})
		whereClause = whereClause.OR(Host.DomainName.IN(domainName...))
	}

	if len(hosts.Alias) > 0 {
		alias := data.Map(hosts.Alias, func(v string) mysql.Expression {
			return mysql.String(v)
		})
		whereClause = whereClause.OR(Host.Alias_.IN(alias...))
	}

	if len(hosts.Both) > 0 {
		data.ForEach(hosts.Both, func(both data.Tuple2[string, string]) {
			criteria1 := Host.DomainName.EQ(mysql.String(both.Left))
			criteria2 := Host.Alias_.EQ(mysql.String(both.Right))
			whereClause = whereClause.OR(criteria1.AND(criteria2))
		})
	}

	statement := Host.UPDATE(Host.IsActive).MODEL(update).WHERE(whereClause)

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
