package core

import (
	"assembler/utils/helper"
	"testing"
)

func TestAssembleJumpLabel(t *testing.T) {
	var tests = []struct {
		param              string
		expectedLabel      *string
		expectedRestOfLine string
		expectsError       bool
	}{
		// Success with label and rest
		{"label:mov 12", helper.StringPointer("label"), "mov 12", false},
		{"label :mov 12", helper.StringPointer("label"), "mov 12", false},
		{"label: mov 12", helper.StringPointer("label"), "mov 12", false},
		{"label : mov 12", helper.StringPointer("label"), "mov 12", false},
		// Success without label
		{"mov 12", nil, "mov 12", false},
		// Success without rest
		{"label:", helper.StringPointer("label"), "", false},
		{"label :", helper.StringPointer("label"), "", false},
		{"label: ", helper.StringPointer("label"), "", false},
		{"label : ", helper.StringPointer("label"), "", false},
		// Fail: Invalid name
		{"1label : mov 12", nil, "mov 12", true},
		{".label :mov 12", nil, "mov 12", true},
		{": mov 12", nil, "mov 12", true},
		// Multiple colons
		{": mov: 12", nil, "mov: 12", true},
		{"label: mov: 12", helper.StringPointer("label"), "mov: 12", false},
	}

	for i, test := range tests {
		gotLabel, gotRestOfLine, err := AssembleJumpLabel(test.param)
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected: %t, Got: %t // param: '%s'", i, test.expectsError, gotErr, test.param)
		}

		if !helper.SafeIsEqualStrPointer(gotLabel, test.expectedLabel) {
			t.Errorf("[%d] Expected: %d, Got: %d // param: '%s'", i, test.expectedLabel, gotLabel, test.param)
		}

		if gotRestOfLine != test.expectedRestOfLine {
			t.Errorf("[%d] Expected: %s, Got: %s // param: '%s'", i, test.expectedRestOfLine, gotRestOfLine, test.param)
		}
	}
}
