package config

import (
	"database/sql"
	"time"

	"mockserver.jiratviriyataranon.io/src/data"
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

type pathToHost struct {
	id        int
	path      string
	hostAlias string
	isActive  bool
	createdAt time.Time
	updatedAt time.Time
}

// DTO
type aliasToHostUpsertMany struct {
	Alias string
	Host  string
}

type pathToHostUpsertMany struct {
	Path      string
	HostAlias string
}

// Request
type registerhostRequest map[string]string

type registerPathRequest map[string]data.StringOrSlice

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
