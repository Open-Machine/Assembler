package utils

import "testing"

func TestStrToInt(t *testing.T) {
	var tests = []struct {
		param        string
		expected     uint
		expectsError bool
	}{
		{"0x10", 16, false},
		{"10", 10, false},
		{"0xa", 10, false},
		{"0xA", 10, false},
		{"0xd", 13, false},
		{"0xg", 0, true},
		{"x10", 0, true},
	}

	for _, test := range tests {
		got, err := StrToPositiveInt(test.param)

		if !(got == test.expected && test.expectsError == (err != nil)) {
			t.Errorf("Expected: %d, Got: %d", test.expected, got)
		}
	}
}

func TestIntToStrHex(t *testing.T) {
	var tests = []struct {
		num        int
		strLength  int
		expected   string
		expectsErr bool
	}{
		{16, 2, "10", false},
		{10, 1, "a", false},
		{17, 2, "11", false},
		{17, 3, "011", false},
		{17, 1, "", true},
	}

	for _, test := range tests {
		got, err := IntToStrHex(test.num, test.strLength)
		gotErr := err != nil

		if !(got == test.expected && gotErr == test.expectsErr) {
			t.Errorf("Wrong")
		}
	}
}
