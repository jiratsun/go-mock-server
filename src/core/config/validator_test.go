package config

import (
	"testing"
)

func TestValidateAuthority(t *testing.T) {
	valid := func(err error) bool {
		return err == nil
	}
	invalid := func(err error) bool {
		return err != nil
	}

	tests := []struct {
		name      string
		authority string
		expected  func(err error) bool
	}{
		{
			"Valid host with user info",
			"https://user:password@www.example.com:443",
			valid,
		},
		{
			"Valid host",
			"https://www.example.com",
			valid,
		},
		{
			"Valid host with port",
			"http://localhost:8080",
			valid,
		},
		{
			"Valid ip address",
			"ftp://192.168.0.1:21",
			valid,
		},
		{
			"Invalid host empty scheme",
			"invalid-url",
			invalid,
		},
		{
			"Invalid host empty host",
			"http://:8080",
			invalid,
		},
		{
			"Invalid host empty protocol",
			"://example.com",
			invalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validateAuthority(test.authority)
			if !test.expected(result) {
				t.Errorf("Test case failed: %v", test.name)
			}
		})
	}
}

func TestValidateAlias(t *testing.T) {
	valid := func(err error) bool {
		return err == nil
	}
	invalid := func(err error) bool {
		return err != nil
	}

	tests := []struct {
		name     string
		alias    string
		expected func(err error) bool
	}{
		{
			"Valid alias",
			"exampleAlias",
			valid,
		},
		{
			"Valid alias with -",
			"example-alias",
			valid,
		},
		{
			"Invalid alias with /",
			"example/alias",
			invalid,
		},
		{
			"Invalid alias with space",
			"example alias",
			invalid,
		},
		{
			"Invalid alias with special character",
			"example!alias",
			invalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validateAlias(test.alias)
			if !test.expected(result) {
				t.Errorf("Test case failed: %v", test.name)
			}
		})
	}
}

func TestValidatePath(t *testing.T) {
	valid := func(err error) bool {
		return err == nil
	}
	invalid := func(err error) bool {
		return err != nil
	}

	tests := []struct {
		name     string
		path     string
		expected func(err error) bool
	}{
		{
			"Valid path",
			"/path/to/resource",
			valid,
		},
		{
			"Valid path with hyphen",
			"/valid-path/",
			valid,
		},
		{
			"Valid path with trailing slash",
			"/another/valid/path/",
			valid,
		},
		{
			"Invalid path with %",
			"/invalid%20path",
			invalid,
		},
		{
			"Invalid path with space",
			"/invalid path",
			invalid,
		},
		{
			"Invalid path with special character",
			"/with_special!characters",
			invalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validatePath(test.path)
			if !test.expected(result) {
				t.Errorf("Test case failed: %v", test.name)
			}
		})
	}
}
