package utils

import "testing"

func TestOverflow(t *testing.T) {
	var tests = []struct {
		num           uint
		availableBits int
		isOverflow    bool
	}{
		{0, 1, false},
		{1, 1, false},
		{2, 1, true},
		{2, 2, false},
		{3, 2, false},
		{4, 2, true},
		{7, 3, false},
		{8, 3, true},
	}

	for _, test := range tests {
		gotIsOverflow := IsOverflow(test.num, test.availableBits)

		if test.isOverflow != gotIsOverflow {
			t.Errorf("Expected overflow: %t, Got overflow: %t // Binary number %d: %b // Available bits: %d", test.isOverflow, gotIsOverflow, test.num, test.num, test.availableBits)
		}
	}
}

func TestIsValidVarName(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"a", true},
		{"var", true},
		{"Var", true},
		{"vaR", true},
		{"VAR", true},
		{"var_able", true},
		{"var_Able", true},
		{"var_ABLE", true},
		{"va0r_4ABLE1", true},
		{"va.", false},
		{"va.r", false},
		{"va-r", false},
		{"va*r", false},
		{"va^r", false},
		{"va&r", false},
		{"&var", false},
		{"var&", false},
		{"jmp", true},
	}

	for _, test := range tests {
		got := IsValidVarName(test.param)

		if test.expected != got {
			t.Errorf("For var name '%s': Expected: %t, Got: %t", test.param, test.expected, got)
		}
	}
}
