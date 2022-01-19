package data

import (
	"reflect"
	"testing"
)

func TestAddInstruction(t *testing.T) {
	program := NewProgram(5)

	if len(program.instructions) != 0 {
		t.Errorf("Expected length 0, got: %d", len(program.instructions))
	}

	program.AddInstruction(Instruction{0, "", 0, 0})

	if len(program.instructions) != 1 {
		t.Errorf("Expected length 1, got: %d", len(program.instructions))
	}
}

func TestAddJumpLabel(t *testing.T) {
	var tests = []struct {
		program      Program
		expectsError bool
	}{
		{Program{[]Instruction{}, map[string]int{"abc": 1, "luca": 2}, map[string]Variable{}}, false},
		{Program{[]Instruction{}, map[string]int{"abc": 1, "label": 2}, map[string]Variable{}}, true},
		{Program{[]Instruction{}, map[string]int{"label": 2}, map[string]Variable{}}, true},
		{Program{[]Instruction{}, map[string]int{"a": 2}, map[string]Variable{}}, false},
	}

	for i, test := range tests {
		err := test.program.AddJumpLabel("label", 1)
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}
	}
}

func TestAddVariable(t *testing.T) {
	program := NewProgram(5)

	if len(program.variablesDict) != 0 {
		t.Errorf("Expected length 0, got: %d", len(program.variablesDict))
	}

	program.AddVariable(Variable{"variable", 0, 3})

	if len(program.variablesDict) != 1 {
		t.Errorf("Expected length 1, got: %d", len(program.variablesDict))
	}
}

func TestProgToExecuter(t *testing.T) {
	tests := []struct {
		instructions []Instruction
		variables    []Variable
		expect       string
		amntErrs     int
	}{
		// Success with header
		{
			[]Instruction{
				{1, "", 2, 0},
				{15, "", 7, 0},
				{0, "", 0, 0},
			},
			[]Variable{
				{
					name:         "a",
					index:        0,
					initialValue: 0x3,
				},
				{
					name:         "c",
					index:        2,
					initialValue: 0x2,
				},
				{
					name:         "b",
					index:        1,
					initialValue: 0x1,
				},
			},
			"v2.0 raw\n0000 1002 f007 0000 4089*0 0002 0001 0003 ",
			0,
		},
		// 1 Overflow
		{
			[]Instruction{
				{1, "", 2, 0},
				{1200, "", 7, 0},
				{0, "", 0, 0},
			},
			[]Variable{},
			"",
			1,
		},
		// 2 Overflows
		{
			[]Instruction{
				{1, "", 2, 0},
				{1200, "", 7, 0},
				{3000, "", 7, 0},
			},
			[]Variable{},
			"",
			2,
		},
	}

	for i, test := range tests {
		program := newProgram(test.instructions, test.variables)
		got, errors := program.ToExecuter()

		if got != test.expect {
			t.Errorf("[%d] Expected: '%s', got: '%s'", i, test.expect, got)
		}

		if len(errors) != test.amntErrs {
			t.Errorf("[%d] Expected %d errors, but got %d", i, len(errors), test.amntErrs)
		}
	}
}
func newProgram(instructions []Instruction, variables []Variable) Program {
	program := NewProgram(len(instructions))
	for _, instruc := range instructions {
		program.AddInstruction(instruc)
	}
	for _, variable := range variables {
		program.AddVariable(variable)
	}
	return program
}

func TestReplaceLabelsWithAddresses(t *testing.T) {
	var tests = []struct {
		programBefore    *Program
		programAfter     Program
		amntErrsExpected int
	}{
		{
			&Program{
				[]Instruction{
					{3, "", 0, NoParam},
					{1, "var", 0, VariableParam},
					{5, "label", 0, JumpLabelParam},
					{7, "var2", 0, VariableParam},
					{7, "label2", 0, JumpLabelParam},
				},
				map[string]int{"label": 3, "label2": 5},
				map[string]Variable{"var": {"var", 1, 4}, "var2": {"var2", 2, 4}},
			},
			Program{
				[]Instruction{
					{3, "", 0, NoParam},
					{1, "var", 4094, VariableParam},
					{5, "label", 3, JumpLabelParam},
					{7, "var2", 4093, VariableParam},
					{7, "label2", 5, JumpLabelParam},
				},
				map[string]int{"label": 3, "label2": 5},
				map[string]Variable{"var": {"var", 1, 4}, "var2": {"var2", 2, 4}},
			},
			0,
		},
		// unused variable and jump label
		{
			&Program{
				[]Instruction{
					{3, "", 0, NoParam},
					{1, "var", 0, VariableParam},
					{5, "label", 0, JumpLabelParam},
					{7, "var2", 0, VariableParam},
					{7, "label2", 0, JumpLabelParam},
				},
				map[string]int{"unused": 1, "label": 3, "label2": 5},
				map[string]Variable{"unusedVar": {"unusedVar", 4, 5}, "var": {"var", 1, 4}, "var2": {"var2", 2, 4}},
			},
			Program{
				[]Instruction{
					{3, "", 0, NoParam},
					{1, "var", 4094, VariableParam},
					{5, "label", 3, JumpLabelParam},
					{7, "var2", 4093, VariableParam},
					{7, "label2", 5, JumpLabelParam},
				},
				map[string]int{"unused": 1, "label": 3, "label2": 5},
				map[string]Variable{"unusedVar": {"unusedVar", 4, 5}, "var": {"var", 1, 4}, "var2": {"var2", 2, 4}},
			},
			0,
		},
		// Fail: multiple jump labels and variables that do not exist
		{
			&Program{
				[]Instruction{
					{3, "", 0, NoParam},
					{1, "var", 0, VariableParam},
					{5, "label", 0, JumpLabelParam},
					{7, "var2", 0, VariableParam},
					{7, "label2", 0, JumpLabelParam},
					{7, "label3", 0, JumpLabelParam},
				},
				map[string]int{"label3": 3},
				map[string]Variable{"var3": {"var3", 1, 4}},
			},
			Program{
				[]Instruction{
					{3, "", 0, NoParam},
					{1, "var", 0, VariableParam},
					{5, "label", 0, JumpLabelParam},
					{7, "var2", 0, VariableParam},
					{7, "label2", 0, JumpLabelParam},
					{7, "label3", 3, JumpLabelParam},
				},
				map[string]int{"label3": 3},
				map[string]Variable{"var3": {"var3", 1, 4}},
			},
			4,
		},
	}

	for i, test := range tests {
		errs := test.programBefore.ReplaceLabelsWithAddresses()

		if test.amntErrsExpected != len(errs) {
			t.Errorf("[%d] Expected %d errors, but got %d", i, test.amntErrsExpected, len(errs))
		}

		if !reflect.DeepEqual(*test.programBefore, test.programAfter) {
			t.Errorf("[%d] Expected program to change to %v, but it changed to %v", i, test.programAfter, *test.programBefore)
		}
	}
}

func TestLenInstructions(t *testing.T) {
	program1 := NewProgram(5)
	if program1.LenInstructions() != 0 {
		t.Errorf("Wrong 1")
	}

	program2 := Program{[]Instruction{mockInstruction(), mockInstruction()}, map[string]int{}, map[string]Variable{}}
	program2.AddInstruction(mockInstruction())
	if program2.LenInstructions() != 3 {
		t.Errorf("Wrong 2")
	}
	program2.AddInstruction(mockInstruction())
	if program2.LenInstructions() != 4 {
		t.Errorf("Wrong 3")
	}
}
func mockInstruction() Instruction {
	return NewInstructionWithoutParam(0)
}
