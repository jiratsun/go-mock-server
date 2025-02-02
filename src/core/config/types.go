package config

import (
	"time"

	"mockserver.jiratviriyataranon.io/src/data"
)

// DB Schema
type hostWithPath struct {
	id             int
	host           string
	alias          string
	isActive       bool
	createdAt      time.Time
	updatedAt      time.Time
	path_id        int
	path_path      string
	path_hostAlias string
	path_isActive  bool
	path_createdAt time.Time
	path_updatedAt time.Time
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
	alias string
	host  string
}

type pathToHostUpsertMany struct {
	path      string
	hostAlias string
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
