package config

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
)

func (request registerHostRequest) valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)
	hosts := request.Hosts

	if hosts == nil {
		problems["hosts"] = "Missing hosts variable"
	}

	if len(hosts) == 0 {
		problems["hosts"] = "Empty hosts variable"
	}

	for _, host := range hosts {
		err := validateAuthority(host.DomainName)
		if err != nil {
			key := fmt.Sprintf("domainName: %v", host.DomainName)
			problems[key] = err.Error()
		}

		err = validateAlias(host.Alias)
		if err != nil {
			key := fmt.Sprintf("alias: %v", host.Alias)
			problems[key] = err.Error()
		}

		if len(host.Description) > 255 {
			key := fmt.Sprintf("description: %v", host.Description)
			problems[key] = "Invalid description: length should not exceed 255"
		}
	}

	return problems
}

func (request modifyHostRequest) valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)
	hosts := request.Hosts

	if hosts == nil {
		problems["hosts"] = "Missing hosts variable"
	}

	if len(hosts) == 0 {
		problems["hosts"] = "Empty hosts variable"
	}

	for _, host := range hosts {
		if host.DomainName != nil {
			err := validateAuthority(*host.DomainName)
			if err != nil {
				key := fmt.Sprintf("domainName: %v", host.DomainName)
				problems[key] = err.Error()
			}
		}

		if host.Alias != nil {
			err := validateAlias(*host.Alias)
			if err != nil {
				key := fmt.Sprintf("alias: %v", host.Alias)
				problems[key] = err.Error()
			}
		}
	}

	return problems
}

func (request registerPathRequest) valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)
	paths := request.Paths

	if paths == nil {
		problems["paths"] = "Missing paths variable"
	}

	if len(paths) == 0 {
		problems["paths"] = "Empty paths variable"
	}

	for _, path := range paths {
		if path.Path == nil {
			problems["path"] = "Missing path variable"
		} else {
			err := validatePath(*path.Path)
			if err != nil {
				key := fmt.Sprintf("path: %v", path.Path)
				problems[key] = err.Error()
			}
		}

		if path.DefaultHost != nil {
			err := validateAuthority(*path.DefaultHost)
			if err != nil {
				key := fmt.Sprintf("defaultHost: %v", path.DefaultHost)
				problems[key] = err.Error()
			}
		}

		if len(path.Description) > 255 {
			key := fmt.Sprintf("description: %v", path.Description)
			problems[key] = "Invalid description: length should not exceed 255"
		}
	}

	return problems
}

func (request modifyPathRequest) valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)
	paths := request.Paths

	if paths == nil {
		problems["paths"] = "Missing paths variable"
	}

	if len(paths) == 0 {
		problems["paths"] = "Empty paths variable"
	}

	for _, path := range paths {
		if path.Path != nil {
			err := validatePath(*path.Path)
			if err != nil {
				key := fmt.Sprintf("path: %v", path.Path)
				problems[key] = err.Error()
			}
		}
	}

	return problems
}

func validateAuthority(authority string) error {
	if authority == "" {
		return errors.New("Invalid authority: empty")
	}

	if len(authority) > 255 {
		return errors.New("Invalid authority: length should not exceed 255")
	}

	parsedAuthority, err := url.Parse(authority)
	if err != nil {
		return fmt.Errorf("Invalid authority: %w", err)
	}

	if parsedAuthority.Scheme == "" {
		return errors.New("Invalid authority: empty scheme")
	}

	host, _, err := net.SplitHostPort(parsedAuthority.Host)
	if err != nil {
		host = parsedAuthority.Host
	}

	if host == "" {
		return errors.New("Invalid authority: empty host")
	}

	return nil
}

func validateAlias(alias string) error {
	if alias == "" {
		return errors.New("Invalid alias: empty")
	}

	if len(alias) > 255 {
		return errors.New("Invalid alias: length should not exceed 255")
	}

	charValidator := regexp.MustCompile(`^([\w\-]*)?$`)
	if !charValidator.MatchString(alias) {
		return errors.New("Invalid alias: invalid characters")
	}

	return nil
}

func validatePath(path string) error {
	if path != "" && !strings.HasPrefix(path, "/") {
		return errors.New("Invalid path: should be empty or begins with /")
	}

	if len(path) > 255 {
		return errors.New("Invalid path: length should not exceed 255")
	}

	charValidator := regexp.MustCompile(`^(/[\w\-./]*)?$`)
	if !charValidator.MatchString(path) {
		return errors.New("Invalid path: invalid characters")
	}

	return nil
}
