package path

import (
	"context"
	"net/url"
)

func (request registerPathRequest) valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	for k, v := range request.PathToHost {
		_, err := url.ParseRequestURI(k + v)
		if err != nil {
			problems[k] = err.Error()
		}
	}

	return problems
}
