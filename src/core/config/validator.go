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

func (request registerhostRequest) valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	for alias, authority := range request {
		key := fmt.Sprintf("%v: %v", alias, authority)

		err := validateAuthority(authority)
		if err != nil {
			problems[key] = err.Error()
		}

		err = validateAlias(alias)
		if err != nil {
			problems[key] = err.Error()
		}
	}

	return problems
}

func (request registerPathRequest) valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	for hostAlias, paths := range request {
		err := validateAlias(hostAlias)
		if err != nil {
			problems[hostAlias] = err.Error()
		}

		for _, path := range paths {
			err = validatePath(path)
			if err != nil {
				key := fmt.Sprintf("%v: %v", hostAlias, path)
				problems[key] = err.Error()
			}
		}
	}

	return problems
}

func validateAuthority(authority string) error {
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

	charValidator := regexp.MustCompile(`^([\w\-]*)?$`)
	if !charValidator.MatchString(alias) {
		return errors.New("Invalid path: invalid characters")
	}

	return nil
}

func validatePath(path string) error {
	if path != "" && !strings.HasPrefix(path, "/") {
		return errors.New("Invalid path: should be empty or begins with /")
	}

	charValidator := regexp.MustCompile(`^(/[\w\-./]*)?$`)
	if !charValidator.MatchString(path) {
		return errors.New("Invalid path: invalid characters")
	}

	return nil
}
