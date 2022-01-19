package core

import (
	"reflect"
	"testing"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/helper"
)

func TestNextAssemblerState(t *testing.T) {
	var tests = []struct {
		param             string
		paramState        int
		expectedNextState int
		isErrExpected     bool
	}{
		// Nothing to change: Success
		{"", CodeState, -1, false},
		{"# Hello World", NoneState, -1, false},
		{"var = 12 # Hello World", VarsState, -1, false},

		// Variable
		{"@VAR", NoneState, VarsState, false},
		{"@VAR", VarsState, -1, true},
		{"@VAR", CodeState, -1, true},

		// Variable
		{"@CODE", VarsState, CodeState, false},
		{"@CODE", NoneState, CodeState, false},
		{"@CODE", CodeState, -1, true},
	}

	for i, test := range tests {
		nextState, err := nextAssemblerState(test.param, test.paramState)

		if test.expectedNextState != nextState {
			t.Errorf("[%d] Expected: %v, Got: %v, Line: '%s'", i, test.expectedNextState, nextState, test.param)
		}

		gotErr := err != nil
		if gotErr != test.isErrExpected {
			t.Errorf("[%d] Error expected: %t, Got error: %t, Error: %v", i, test.isErrExpected, gotErr, err)
		}
	}
}

func TestAssembleState(t *testing.T) {
	var tests = []struct {
		param         string
		paramState    int
		expectedState int
		isErrExpected bool
	}{
		// Blank: Success
		{"", CodeState, CodeState, false},
		{"	 	 	", VarsState, VarsState, false},
		{"# Hello World", NoneState, NoneState, false},

		// Variable
		{"@VAR #comment", NoneState, VarsState, false},
		{"@VAR #comment", VarsState, VarsState, true},
		{"@VAR", CodeState, CodeState, true},

		// Variable
		{"@CODE", VarsState, CodeState, false},
		{"@CODE", NoneState, CodeState, false},
		{"@CODE#comment", CodeState, CodeState, true},
	}

	for i, test := range tests {
		variableIndex := 0
		stateParam := test.paramState
		gotLabel, gotVariable, gotCmd, err := assembleEntireLine(test.param, &stateParam, &variableIndex)

		if test.expectedState != stateParam {
			t.Errorf("[%d] Expected state: %v | Got: %v | Line: '%s'", i, test.expectedState, stateParam, test.param)
		}

		if gotLabel != nil || gotCmd != nil || gotVariable != nil {
			t.Errorf("[%d] Jump label and command should be null. Got: %v | %v | %v", i, gotLabel, gotCmd, gotVariable)
		}

		gotErr := err != nil
		if gotErr != test.isErrExpected {
			t.Errorf("[%d] Error expected: %t, Got error: %t, Error: %v", i, test.isErrExpected, gotErr, err)
		}
	}
}

func TestAssembleVariableEntireLine(t *testing.T) {
	var tests = []struct {
		param            string
		expectedVariable *data.Variable
		isErrExpected    bool
	}{
		// Blank: Success
		{"", nil, false},
		{"	 	 	", nil, false},
		{"# Hello World", nil, false},

		// With comment: Success
		{"variable=5#Hello World ", newVariable("variable", 5), false},

		// Different spaces: Success
		{"	a  =   1", newVariable("a", 1), false},
		{"	a=0x1 ", newVariable("a", 1), false},
		{"a=0x1", newVariable("a", 1), false},
		{"a= 0x1", newVariable("a", 1), false},
		{"a =0x1", newVariable("a", 1), false},

		// Different numbers: Success
		{"a = 0x1", newVariable("a", 1), false},
		{"a = 1", newVariable("a", 1), false},

		// Too large number: Fail
		{"a = 200000", nil, true},

		// Nothing: Fail (as instruction)
		{"	; ", nil, true},

		// Reserved word: Fail
		{"copy = 4", nil, true},
	}

	for i, test := range tests {
		state := VarsState
		variableIndex := 0
		gotLabel, gotVariable, gotCmd, err := assembleEntireLine(test.param, &state, &variableIndex)

		if state != VarsState {
			t.Errorf("[%d] Shouldn't have changed the state", i)
		}

		if !reflect.DeepEqual(gotVariable, test.expectedVariable) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedVariable, gotVariable)
		}

		if gotVariable == nil {
			if variableIndex != 0 {
				t.Errorf("[%d] Shouldn't have increased variableIndex", i)
			}
		} else {
			if variableIndex != 1 {
				t.Errorf("[%d] Should have increased variableIndex to 1", i)
			}
		}

		if gotLabel != nil || gotCmd != nil {
			t.Errorf("[%d] Jump label and command should be null. Got: %v | %v", i, gotLabel, gotCmd)
		}

		gotErr := err != nil
		if gotErr != test.isErrExpected {
			t.Errorf("[%d] Error expected: %t, Error expected: %t", i, test.isErrExpected, gotErr)
		}
	}
}
func newVariable(name string, initialValue uint) *data.Variable {
	value, _ := data.NewVariable(name, 0, initialValue)
	return value
}

func TestAssembleCodeEntireLine(t *testing.T) {
	config.Testing = true

	var tests = []struct {
		param         string
		expectedLabel *string
		expectedCmd   *data.Instruction
		isErrExpected bool
	}{
		// Blank: Success
		{"", nil, nil, false},
		{"	 	 	", nil, nil, false},

		// With comment
		{"# Hello World", nil, nil, false},
		{"input var#Hello World ", nil, newInstructionStr(7, "var"), false},
		{"input a  #Hello World ", nil, newInstructionStr(7, "a"), false},
		{"hello:#Hello World ", helper.StringPointer("hello"), nil, false},

		// Number of parameters: Fail
		{"input #1 ", nil, nil, true},

		// Instruction: Success
		{"	input 	var ", nil, newInstructionStr(7, "var"), false},
		// Instruction: Fail
		{"	input 	0x1 ", nil, nil, true},
		{"	inputa 	a ", nil, nil, true},

		// Label: Success
		{"	label: ", helper.StringPointer("label"), nil, false},
		// Label: Fail
		{"	1label: input 	0x1 ", nil, nil, true},

		// More than one: Fail
		{"	label: input 	a ", nil, nil, true},
		{"	label: input b ", nil, nil, true},

		// Nothing: Fail (as instruction)
		{"	; ", nil, nil, true},

		// Reserved word: Fail
		{"copy: ", nil, nil, true},
	}

	for i, test := range tests {
		state := CodeState
		variableIndex := 0
		gotLabel, gotVariable, gotCmd, err := assembleEntireLine(test.param, &state, &variableIndex)

		if state != CodeState {
			t.Errorf("[%d] Shouldn't have changed the state", i)
		}
		if gotVariable != nil {
			t.Errorf("[%d] Variable should be null: %v, Got: %v", i, test.expectedCmd, gotCmd)
		}

		if !helper.SafeIsEqualInstructionPointer(test.expectedCmd, gotCmd) {
			t.Errorf("[%d] Expected cmd: %v, Got: %v, Line: '%s'", i, test.expectedCmd, gotCmd, test.param)
		}

		if !helper.SafeIsEqualStrPointer(gotLabel, test.expectedLabel) {
			t.Errorf("[%d] Expected label: %v, Got: %v", i, test.expectedLabel, gotLabel)
		}

		gotErr := err != nil
		if gotErr != test.isErrExpected {
			t.Errorf("[%d] Error expected: %t, Got error: %t, Error: %v, Line: '%s'", i, test.isErrExpected, gotErr, err, test.param)
		}
	}
}
func newInstructionStr(cmdCode int, param string) *data.Instruction {
	return data.NewInstructionTest(cmdCode, param, 0, 0)
}
