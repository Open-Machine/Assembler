package utils

import (
	"testing"
)

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
		expected int
	}{
		// cammelCase
		{"var", validName},
		{"vaR", validName},
		{"varName", validName},
		// [FAIL] snake_case
		{"var_name", invalidNameRegex},
		// [FAIL] PascalCase
		{"Var", invalidNameRegex},
		// [FAIL] ALLUPERCASE
		{"VAR", invalidNameRegex},
		// [FAIL] Special*Characters
		{"va.", invalidNameRegex},
		{"va.r", invalidNameRegex},
		{"va-r", invalidNameRegex},
		{"va*r", invalidNameRegex},
		{"va^r", invalidNameRegex},
		{"va&r", invalidNameRegex},
		{"&var", invalidNameRegex},
		{"var&", invalidNameRegex},
		// [FAIL] Blank
		{"", invalidNameRegex},
		// Reserved words
		{"jmp", reservedWord},
		{"copy", reservedWord},
	}

	for _, test := range tests {
		got := nameStatus(test.param)

		if test.expected != got {
			t.Errorf("For var name '%s': Expected: %d, Got: %d", test.param, test.expected, got)
		}

		gotError := got != validName

		gotParam := CheckParamName(test.param)
		gotErrorParam := gotParam != nil

		gotJump := CheckJumpLabelName(test.param)
		gotErrorJump := gotJump != nil

		if !(gotError == gotErrorParam && gotError == gotErrorJump) {
			t.Errorf("CheckParamName and CheckJumpLabelName have is error returns. General got error: %t, Jump got error: %t, Param got error: %t", gotError, gotErrorJump, gotErrorParam)
		}
	}
}

func TestIsReservedWord(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		// Not reserved
		{"mov", false},
		{"variable", false},
		{"name", false},
		// Reserved
		{"jmp", true},
		{"store", true},
		// Maybe in future
		{"procedure", false},
		{"import", false},
		{"declare", false},
		{"var", false},
	}

	for _, test := range tests {
		got := isReservedWord(test.param)

		if test.expected != got {
			t.Errorf("For var name '%s': Expected: %t, Got: %t", test.param, test.expected, got)
		}
	}
}
