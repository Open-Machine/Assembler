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
	cmd, _ := data.NewCommand(a, data.NewIntParam(b))
	program := data.ProgramFromCommandsAndLabels([]data.Command{*cmd}, map[string]int{})
	return &program
}

func TestSafeIsEqualCommandPointer(t *testing.T) {
	var tests = []struct {
		param1   *data.Command
		param2   *data.Command
		expected bool
	}{
		{nil, nil, true},
		{newCommand(1, 1), nil, false},
		{nil, newCommand(1, 1), false},
		{newCommand(1, 1), newCommand(1, 1), true},
		{newCommand(1, 1), newCommand(1, 2), false},
	}

	for i, test := range tests {
		got := SafeIsEqualCommandPointer(test.param1, test.param2)
		if got != test.expected {
			t.Errorf("[%d] Expected: %t, Got: %t", i, test.expected, got)
		}
	}
}
func newCommand(a int, b int) *data.Command {
	cmd, _ := data.NewCommand(a, data.NewIntParam(b))
	return cmd
}

func TestSafeIsEqualCommandParamPointer(t *testing.T) {
	var tests = []struct {
		param1   *data.CommandParameter
		param2   *data.CommandParameter
		expected bool
	}{
		{nil, nil, true},
		{newIntParam(1), nil, false},
		{nil, newIntParam(1), false},
		{newIntParam(1), newIntParam(1), true},
		{newIntParam(1), newIntParam(2), false},
	}

	for i, test := range tests {
		got := SafeIsEqualCommandParamPointer(test.param1, test.param2)
		if got != test.expected {
			t.Errorf("[%d] Expected: %t, Got: %t", i, test.expected, got)
		}
	}
}
func newIntParam(a int) *data.CommandParameter {
	param := data.NewIntParam(a)
	return &param
}
