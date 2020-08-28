package core

import (
	"testing"

	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/utils/helper"
)

func TestAssembleEntireLine(t *testing.T) {
	var tests = []struct {
		param         string
		expectedLabel *string
		expectedCmd   *data.Instruction
		isErrExpected bool
	}{
		// Blank: Success
		{"", nil, nil, false},
		{"	 	 	", nil, nil, false},

		// Instruction: Success
		{"	input 	1 ", nil, newInstruction(7, 1), false},
		{"	input 	0x1 ", nil, newInstruction(7, 1), false},
		// Instruction: Fail
		{"	inputa 	0x1 ", nil, nil, true},
		{"	input 	a0x1 ", nil, nil, true},

		// Label: Success
		{"	label: ", helper.StringPointer("label"), nil, false},
		// Label: Fail
		{"	1label: input 	0x1 ", nil, nil, true},

		// More than one: Fail
		{"	label: input 	0x1 ", nil, nil, true},
		{"	1label: input 0x1 ", nil, nil, true},

		// Nothing: Fail (as instruction)
		{"	; ", nil, nil, true},
	}

	for i, test := range tests {
		gotLabel, gotCmd, err := assembleEntireLine(test.param)

		if !helper.SafeIsEqualInstructionPointer(test.expectedCmd, gotCmd) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedCmd, gotCmd)
		}

		if !helper.SafeIsEqualStrPointer(gotLabel, test.expectedLabel) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedLabel, gotLabel)
		}

		gotErr := err != nil
		if gotErr != test.isErrExpected {
			t.Errorf("[%d] Error expected: %t, Error expected: %t", i, test.isErrExpected, gotErr)
		}
	}
}
