package path

import "database/sql"

type PathStore struct {
	SqlPool *sql.DB
}
