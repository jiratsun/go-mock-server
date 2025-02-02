package host

import (
	"context"
	"testing"
)

func TestRegisterHostRequestValid(t *testing.T) {
	valid := func(problems map[string]string) bool {
		return len(problems) == 0
	}
	invalid := func(problems map[string]string) bool {
		return len(problems) > 0
	}

	tests := []struct {
		name     string
		request  registerhostRequest
		expected func(map[string]string) bool
	}{
		{
			"Valid host with user info",
			registerhostRequest{"alias": "https://user:password@www.example.com:443"},
			valid,
		},
		{
			"Valid host",
			registerhostRequest{"alias": "https://www.example.com"},
			valid,
		},
		{
			"Valid host with port",
			registerhostRequest{"alias": "http://localhost:8080"},
			valid,
		},
		{
			"Valid ip address",
			registerhostRequest{"alias": "ftp://192.168.0.1:21"},
			valid,
		},
		{
			"Invalid host empty scheme",
			registerhostRequest{"alias": "invalid-url"},
			invalid,
		},
		{
			"Invalid host empty host",
			registerhostRequest{"alias": "http://:8080"},
			invalid,
		},
		{
			"Invalid host empty protocol",
			registerhostRequest{"alias": "://example.com"},
			invalid,
		},
		{
			"Valid alias",
			registerhostRequest{"exampleAlias": "https://www.host.com"},
			valid,
		},
		{
			"Valid alias with -",
			registerhostRequest{"example-alias": "https://www.host.com"},
			valid,
		},
		{
			"Invalid alias with /",
			registerhostRequest{"example/alias": "https://www.host.com"},
			invalid,
		},
		{
			"Invalid alias with space",
			registerhostRequest{"example alias": "https://www.host.com"},
			invalid,
		},
		{
			"Invalid alias with special character",
			registerhostRequest{"example!alias": "https://www.host.com"},
			invalid,
		},
	}

	for _, test := range tests {
		result := test.request.valid(context.Background())
		if !test.expected(result) {
			t.Errorf("Test case failed: %v", test.name)
		}
	}
}
