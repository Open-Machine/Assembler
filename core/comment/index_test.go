package comment

import "testing"

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
		got := RemoveComment(test.param)
		if got != test.expected {
			t.Errorf("Comment removed wrongly. Expected '%s', Got: '%s'", test.expected, test.param)
		}
	}
}
