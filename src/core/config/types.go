package config

import (
	"database/sql"
	"time"
)

// DB Schema
type hostWithPath struct {
	id            int
	host          string
	alias         string
	isActive      bool
	createdAt     time.Time
	updatedAt     time.Time
	pathId        sql.NullInt64
	path          sql.NullString
	hostAlias     sql.NullString
	pathIsActive  sql.NullBool
	pathCreatedAt sql.NullTime
	pathUpdatedAt sql.NullTime
}

type path struct {
	id        int
	path      string
	hostAlias string
	isActive  bool
	createdAt time.Time
	updatedAt time.Time
}

// DTO
type hostUpsertMany struct {
	DomainName  string
	Alias       string
	Description string
}

type pathUpsertMany struct {
	Path        string
	DefaultHost sql.NullString
	Description string
}

// Request
type registerhostRequest struct {
	Hosts []struct {
		DomainName  string
		Alias       string
		Description string
	}
}

type registerPathRequest struct {
	Paths []struct {
		Path        string
		DefaultHost string
		Description string
	}
}

// Response
type getHostResponse map[string]*hostInfo

type hostInfo struct {
	Host     string
	IsActive bool
	Paths    []pathInfo
}

type pathInfo struct {
	Path     string
	IsActive bool
}
