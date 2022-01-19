package helper

import (
	"testing"

	"github.com/open-machine/assembler/data"
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
		{newProgram(1, "a"), nil, false},
		{nil, newProgram(1, "a"), false},
		{newProgram(1, "a"), newProgram(1, "a"), true},
		{newProgram(1, "a"), newProgram(1, "b"), false},
	}

	for i, test := range tests {
		got := SafeIsEqualProgramPointer(test.param1, test.param2)
		if got != test.expected {
			t.Errorf("[%d] Expected: %t, Got: %t", i, test.expected, got)
		}
	}
}

func newProgram(a int, b string) *data.Program {
	cmd, _ := data.NewJumpInstruction(a, b)
	program := data.NewCompleteProgram([]data.Instruction{*cmd}, map[string]int{}, map[string]data.Variable{})
	return &program
}

func TestSafeIsEqualInstructionPointer(t *testing.T) {
	var tests = []struct {
		param1   *data.Instruction
		param2   *data.Instruction
		expected bool
	}{
		{nil, nil, true},
		{newInstruction(1, "a"), nil, false},
		{nil, newInstruction(1, "a"), false},
		{newInstruction(1, "a"), newInstruction(1, "a"), true},
		{newInstruction(1, "a"), newInstruction(1, "b"), false},
	}

	for i, test := range tests {
		got := SafeIsEqualInstructionPointer(test.param1, test.param2)
		if got != test.expected {
			t.Errorf("[%d] Expected: %t, Got: %t", i, test.expected, got)
		}
	}
}
func newInstruction(a int, b string) *data.Instruction {
	cmd, _ := data.NewVariableInstruction(a, b)
	return cmd
}
