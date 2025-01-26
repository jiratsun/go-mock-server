package path

import "time"

type pathToHost struct {
	id        int
	path      string
	host      string
	isActive  bool
	createdAt time.Time
	updatedAt time.Time
}

type pathToHostUpsertMany struct {
	path string
	host string
}

type registerPathRequest map[string]string

type getPathResponse map[string]getPathInfo

type getPathInfo struct {
	Host     string
	IsActive bool
}
