package utils

import "testing"

func TestLineNormalization(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"  MOV\t 8 \r\n", "MOV 8"},
		{"MoV 8\n", "MoV 8"},
		{"MoV 8 #asdf", "MoV 8"},
		{"hello, my name is Luca #adfd \n", "hello, my name is Luca"},
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
		got := removeNewLine(test.param)

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

func TestRemoveComment(t *testing.T) {
	tests := []struct {
		param    string
		expected string
	}{
		{"", ""},
		{"copy 0x0", "copy 0x0"},
		{"# Hello World ", ""},
		{"copy 0x0 # Hello World", "copy 0x0 "},
		{"luca dillenburg doing stuff # Hello World", "luca dillenburg doing stuff "},
	}

	for _, test := range tests {
		got := removeComment(test.param)
		if got != test.expected {
			t.Errorf("Comment removed wrongly. Expected '%s', Got: '%s'", test.expected, test.param)
		}
	}
}
