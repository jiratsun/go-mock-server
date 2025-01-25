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

type pathRequest struct {
	PathToHost map[string]string
}
