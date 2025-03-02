package config

import (
	"database/sql"
)

// DTO
type hostUpsertMany struct {
	DomainName  string
	Alias       string
	Description string
}

type pathUpsertMany struct {
	Path        string
	DefaultHost sql.NullString
	Description string
}

// Request
type registerhostRequest struct {
	Hosts []struct {
		DomainName  string
		Alias       string
		Description string
	}
}

type registerPathRequest struct {
	Paths []struct {
		Path        string
		DefaultHost string
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

type pathInfo struct {
	Path     string
	IsActive bool
}
