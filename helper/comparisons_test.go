package helper

import (
	"github.com/open-machine/assembler/data"
	"testing"
)

func TestSafeIsEqualStringPointer(t *testing.T) {
	var tests = []struct {
		param1   *string
		param2   *string
		expected bool
	}{
		{nil, nil, true},
		{StringPointer("Hello"), nil, false},
		{nil, StringPointer("Hello"), false},
		{StringPointer("Hello"), StringPointer("Hello"), true},
		{StringPointer("Hello"), StringPointer("Hella"), false},
	}

	for i, test := range tests {
		got := SafeIsEqualStrPointer(test.param1, test.param2)
		if got != test.expected {
			t.Errorf("[%d] Expected: %t, Got: %t", i, test.expected, got)
		}
	}
}

func TestSafeIsEqualProgramPointer(t *testing.T) {
	var tests = []struct {
		param1   *data.Program
		param2   *data.Program
		expected bool
	}{
		{nil, nil, true},
		{newProgram(1, 1), nil, false},
		{nil, newProgram(1, 1), false},
		{newProgram(1, 1), newProgram(1, 1), true},
		{newProgram(1, 1), newProgram(1, 2), false},
	}

	for i, test := range tests {
		got := SafeIsEqualProgramPointer(test.param1, test.param2)
		if got != test.expected {
			t.Errorf("[%d] Expected: %t, Got: %t", i, test.expected, got)
		}
	}
}
func newProgram(a int, b int) *data.Program {
	cmd, _ := data.NewInstruction(a, data.NewIntParam(b))
	program := data.ProgramFromInstructionsAndLabels([]data.Instruction{*cmd}, map[string]int{})
	return &program
}

func TestSafeIsEqualInstructionPointer(t *testing.T) {
	var tests = []struct {
		param1   *data.Instruction
		param2   *data.Instruction
		expected bool
	}{
		{nil, nil, true},
		{newInstruction(1, 1), nil, false},
		{nil, newInstruction(1, 1), false},
		{newInstruction(1, 1), newInstruction(1, 1), true},
		{newInstruction(1, 1), newInstruction(1, 2), false},
	}

	for i, test := range tests {
		got := SafeIsEqualInstructionPointer(test.param1, test.param2)
		if got != test.expected {
			t.Errorf("[%d] Expected: %t, Got: %t", i, test.expected, got)
		}
	}
}
func newInstruction(a int, b int) *data.Instruction {
	cmd, _ := data.NewInstruction(a, data.NewIntParam(b))
	return cmd
}

func TestSafeIsEqualInstructionParamPointer(t *testing.T) {
	var tests = []struct {
		param1   *data.InstructionParameter
		param2   *data.InstructionParameter
		expected bool
	}{
		{nil, nil, true},
		{newIntParam(1), nil, false},
		{nil, newIntParam(1), false},
		{newIntParam(1), newIntParam(1), true},
		{newIntParam(1), newIntParam(2), false},
	}

	for i, test := range tests {
		got := SafeIsEqualInstructionParamPointer(test.param1, test.param2)
		if got != test.expected {
			t.Errorf("[%d] Expected: %t, Got: %t", i, test.expected, got)
		}
	}
}
func newIntParam(a int) *data.InstructionParameter {
	param := data.NewIntParam(a)
	return &param
}
