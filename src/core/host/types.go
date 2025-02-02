package host

import "time"

type host struct {
	id        int
	host      string
	alias     string
	isActive  bool
	createdAt time.Time
	updatedAt time.Time
}

type aliasToHostUpsertMany struct {
	alias string
	host  string
}

type registerhostRequest map[string]string

type getHostResponse map[string]hostInfo

type hostInfo struct {
	Host     string
	IsActive bool
}
