package utils

import "testing"

func TestLineNormalization(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"  MOV\t 8 \r\n", "mov 8"},
		{"MoV 8\n", "mov 8"},
	}

	for _, test := range tests {
		got := LineNormalization(test.param)
		if got != test.expected {
			t.Errorf("Expected: '%s', Got: '%s'", test.expected, got)
		}
	}
}

func TestRemoveNewLine(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		// Unix
		{"mov 8\n", "mov 8"},
		{"mov 8 \n", "mov 8 "},
		// Windows
		{"mov acc 8\r\n", "mov acc 8"},
		{"mov acc 8 \r\n", "mov acc 8 "},
	}

	for _, test := range tests {
		got := RemoveNewLine(test.param)

		if got != test.expected {
			t.Errorf("Expected: '%s', Got: '%s'", test.expected, got)
		}
	}
}

func TestRemoveUnecessarySpaces(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"\t mov 8\t ", "mov 8"},
		{"mov\tacc \t 	8", "mov acc 8"},
		{"\tmov \tacc \t	8", "mov acc 8"},
	}

	for _, test := range tests {
		got := removeUnecessarySpaces(test.param)

		if got != test.expected {
			t.Errorf("Expected: '%s', Got: '%s'", test.expected, got)
		}
	}
}
