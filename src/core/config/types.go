package config

import (
	"database/sql"

	"mockserver.jiratviriyataranon.io/src/data"
)

// DTO
type hostUpsertMany struct {
	DomainName  string
	Alias       string
	Description string
}

type hostDeleteMany struct {
	DomainName []string
	Alias      []string
	Both       []data.Tuple2[string, string]
}

type pathUpsertMany struct {
	Path        string
	DefaultHost sql.NullString
	Description string
}

// Request
type registerHostRequest struct {
	Hosts []struct {
		DomainName  string
		Alias       string
		Description string
	}
}

type deleteHostRequest struct {
	Hosts []struct {
		DomainName *string
		Alias      *string
	}
}

type registerPathRequest struct {
	Paths []struct {
		Path        string
		DefaultHost *string
		Description string
	}
}

// Response
type getHostResponse map[string]*hostInfo

type hostInfo struct {
	Alias       string `json:"alias"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

type getPathResponse map[string]*pathInfo

type pathInfo struct {
	DefaultHost *string
	Description string
	IsActive    bool
}
