package core

import (
	"testing"

	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/utils/helper"
)

func TestAssembleEntireLine(t *testing.T) {
	var tests = []struct {
		param            string
		expectedLabel    *string
		expectedCmd      *data.Instruction
		amntErrsExpected int
	}{
		// Success without label
		{"", nil, nil, 0},
		{" 	 ", nil, nil, 0},
		{"	 	 	", nil, nil, 0},
		{"	input 	1 ", nil, newInstruction(7, 1), 0},
		{"	input 	0x1 ", nil, newInstruction(7, 1), 0},
		// Fail without label
		{"	inputa 	0x1 ", nil, nil, 1},
		{"	inputa 	a0x1 ", nil, nil, 1},
		// Success with label
		{"	label: input 	0x1 ", helper.StringPointer("label"), newInstruction(7, 1), 0},
		// Success only label
		{"	label: ", helper.StringPointer("label"), nil, 0},
		// Fail with label
		{"	label: inputa 	0x1 ", helper.StringPointer("label"), nil, 1},
		{"	1label: inputa 	0x1 ", nil, nil, 2},
		{"	1label: input 	0x1 ", nil, newInstruction(7, 1), 1},
	}

	for i, test := range tests {
		gotLabel, gotCmd, errs := assembleEntireLine(test.param)

		if !helper.SafeIsEqualInstructionPointer(test.expectedCmd, gotCmd) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedCmd, gotCmd)
		}

		if !helper.SafeIsEqualStrPointer(gotLabel, test.expectedLabel) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedLabel, gotLabel)
		}

		if len(errs) != test.amntErrsExpected {
			t.Errorf("[%d] Amnt errors expected: %d, Errors: %v", i, test.amntErrsExpected, errs)
		}
	}
}
