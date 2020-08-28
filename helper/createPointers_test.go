package helper

import (
	"testing"
)

func TestStringPointer(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{""},
		{"a"},
		{"abc"},
		{"aBc"},
	}

	for _, test := range tests {
		got := StringPointer(test.param)
		if *got != test.param {
			t.Errorf("Expected: '%s', Got: '%s'", test.param, *got)
		}
	}
}
