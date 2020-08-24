package helper

import "testing"

func TestFilenameWithoutExtension(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"file", "file"},
		{"fileA", "fileA"},
		{"FileA", "FileA"},
		{"FileA.txt", "FileA"},
		{"FileA.hello", "FileA"},
		{"aaa.bbb", "aaa"},
	}

	for _, test := range tests {
		got := FilenameWithoutExtension(test.param)
		if got != test.expected {
			t.Errorf("Expected: '%s', Got: '%s'", test.expected, got)
		}
	}
}
