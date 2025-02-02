package path

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"mockserver.jiratviriyataranon.io/src/core/host"
)

func (request registerPathRequest) valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	for hostAlias, paths := range request {
		err := host.ValidateAlias(hostAlias)
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
