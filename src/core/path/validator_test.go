package path

import (
	"context"
	"testing"
)

func TestRegisterPathRequestValid(t *testing.T) {
	valid := func(problems map[string]string) bool {
		return len(problems) == 0
	}
	invalid := func(problems map[string]string) bool {
		return len(problems) > 0
	}

	tests := []struct {
		name     string
		request  registerPathRequest
		expected func(map[string]string) bool
	}{
		{
			"Valid host with user info",
			registerPathRequest{"/validPath": "https://user:password@www.example.com:443"},
			valid,
		},
		{
			"Valid host",
			registerPathRequest{"/validPath": "https://www.example.com"},
			valid,
		},
		{
			"Valid host with port",
			registerPathRequest{"/validPath": "http://localhost:8080"},
			valid,
		},
		{
			"Valid ip address",
			registerPathRequest{"/validPath": "ftp://192.168.0.1:21"},
			valid,
		},
		{
			"Invalid host empty scheme",
			registerPathRequest{"/validPath": "invalid-url"},
			invalid,
		},
		{
			"Invalid host empty host",
			registerPathRequest{"/validPath": "http://:8080"},
			invalid,
		},
		{
			"Invalid host empty protocol",
			registerPathRequest{"/validPath": "://example.com"},
			invalid,
		},
		{
			"Valid path",
			registerPathRequest{"/path/to/resource": "https://www.validHost.com"},
			valid,
		},
		{
			"Valid path with -",
			registerPathRequest{"/valid-path/": "https://www.validHost.com"},
			valid,
		},
		{
			"Valid path with trailing /",
			registerPathRequest{"/another/valid/path/": "https://www.validHost.com"},
			valid,
		},
		{
			"Invalid path with %",
			registerPathRequest{"/invalid%20path": "https://www.validHost.com"},
			invalid,
		},
		{
			"Invalid path with space",
			registerPathRequest{"/invalid path": "https://www.validHost.com"},
			invalid,
		},
		{
			"Invalid path with special character",
			registerPathRequest{"/with_special!characters": "https://www.validHost.com"},
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
