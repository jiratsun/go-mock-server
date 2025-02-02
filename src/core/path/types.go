package path

import "time"

type pathToHost struct {
	id         int
	path       string
	host_alias string
	isActive   bool
	createdAt  time.Time
	updatedAt  time.Time
}

type pathToHostUpsertMany struct {
	path       string
	host_alias string
}

type registerPathRequest map[string]string

type getPathResponse map[string]pathInfo

type pathInfo struct {
	Host     string
	IsActive bool
}
