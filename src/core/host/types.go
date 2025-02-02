package host

import (
	"time"
)

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

type aliasToHostUpsertMany struct {
	alias string
	host  string
}

type registerhostRequest map[string]string

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
