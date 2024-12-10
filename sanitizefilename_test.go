package sanitizefilename

import (
	"testing"
)

func TestSanitizeWindows(t *testing.T) {
	simulation = "windows"
	defer func() {
		simulation = ""
	}()
	tests := []struct {
		input    string
		expected string
	}{
		{"   ", ""},
		{"example", "example"},
		{"abc.txt", "abc.txt"},
		{"<>:\"/\\|?*abc.txt", "abc.txt"},
		{"abc\x1f.txt", "abc.txt"},
		{"abc\x7f.txt", "abc.txt"},
		{"NUL", ""},
		{"NUL.txt", ".txt"},
		{"NUL.tar.gz", ".tar.gz"},
	}

	for _, test := range tests {
		result := Sanitize(test.input)
		if result != test.expected {
			t.Errorf("Sanitize(\"%s\") = \"%s\"; expected \"%s\"", test.input, result, test.expected)
		}
	}
}

func TestSanitizeLinuxAndUnix(t *testing.T) {
	simulation = "not windows"
	defer func() {
		simulation = ""
	}()
	tests := []struct {
		input    string
		expected string
	}{
		{"   ", "   "},
		{"example", "example"},
		{"abc.txt", "abc.txt"},
		{"abc\x1f.txt", "abc.txt"},
		{"abc\x7f.txt", "abc.txt"},
		{".", ""},
		{"..", ""},
		{"/..", ""},
	}

	for _, test := range tests {
		result := Sanitize(test.input)
		if result != test.expected {
			t.Errorf("Sanitize(\"%s\") = \"%s\"; expected \"%s\"", test.input, result, test.expected)
		}
	}
}
