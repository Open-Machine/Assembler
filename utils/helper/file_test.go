package helper

import "testing"

func TestFileNameWithoutExtension(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"FileA", "FileA"},
		{"directory/file", "directory/file"},
		{"directory/file.txt", "directory/file"},
		{"/Users/name/file.sbt", "/Users/name/file"},
		{"FileA.txt", "FileA"},
		{"FileA.hello", "FileA"},
		{"aaa.bbb", "aaa"},
	}

	for _, test := range tests {
		got := FileNameWithoutExtension(test.param)
		if got != test.expected {
			t.Errorf("Expected: '%s', Got: '%s'", test.expected, got)
		}
	}
}

func TestFileNameExtension(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"file.asm", ".asm"},
		{"path/a.txt", ".txt"},
		{"file", ""},
	}

	for _, test := range tests {
		got := FileExtension(test.param)
		if got != test.expected {
			t.Errorf("Expected: '%s', Got: '%s'", test.expected, got)
		}
	}
}
