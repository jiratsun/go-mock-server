package path

import (
	"time"

	"mockserver.jiratviriyataranon.io/src/data"
)

type pathToHost struct {
	id        int
	path      string
	hostAlias string
	isActive  bool
	createdAt time.Time
	updatedAt time.Time
}

type pathToHostUpsertMany struct {
	path      string
	hostAlias string
}

type registerPathRequest map[string]data.StringOrSlice
