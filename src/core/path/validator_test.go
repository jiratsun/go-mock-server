package path

import (
	"testing"
)

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
