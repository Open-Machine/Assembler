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

	program.AddInstruction(Instruction{0, NewIntParam(0)})

	if len(program.instructions) != 1 {
		t.Errorf("Expected length 1, got: %d", len(program.instructions))
	}
}

func TestProgToExecuter(t *testing.T) {
	tests := []struct {
		instructions []Instruction
		expect       string
		amntErrs     int
	}{
		// Success with header
		{
			[]Instruction{
				{1, NewIntParam(2)},
				{15, NewIntParam(7)},
				{0, NewIntParam(0)},
			},
			"v2.0 raw\n0000 1002 f007 0000",
			0,
		},
		// 1 Overflow
		{
			[]Instruction{
				{1, NewIntParam(2)},
				{1200, NewIntParam(7)},
				{0, NewIntParam(0)},
			},
			"",
			1,
		},
		// 2 Overflows
		{
			[]Instruction{
				{1, NewIntParam(2)},
				{1200, NewIntParam(7)},
				{3000, NewIntParam(7)},
			},
			"",
			2,
		},
	}

	for i, test := range tests {
		program := newProgram(test.instructions)
		got, errors := program.ToExecuter()

		if got != test.expect {
			t.Errorf("[%d] Expected: '%s', got: '%s'", i, test.expect, got)
		}

		if len(errors) != test.amntErrs {
			t.Errorf("[%d] Expected %d errors, but got %d", i, len(errors), test.amntErrs)
		}
	}
}
func newProgram(instructions []Instruction) Program {
	program := NewProgram(len(instructions))
	for _, instruc := range instructions {
		program.AddInstruction(instruc)
	}
	return program
}

func TestAddJumpLabel(t *testing.T) {
	var tests = []struct {
		program      Program
		expectsError bool
	}{
		{Program{[]Instruction{}, map[string]int{"abc": 1, "luca": 2}}, false},
		{Program{[]Instruction{}, map[string]int{"abc": 1, "label": 2}}, true},
		{Program{[]Instruction{}, map[string]int{"label": 2}}, true},
		{Program{[]Instruction{}, map[string]int{"a": 2}}, false},
	}

	for i, test := range tests {
		err := test.program.AddJumpLabel("label", 1)
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}
	}
}

func TestReplaceLabelsWithNumbers(t *testing.T) {
	var tests = []struct {
		programBefore    *Program
		programAfter     Program
		amntErrsExpected int
	}{
		// single jump label
		{
			&Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewStringParam("label")},
					{5, NewIntParam(3)},
					{7, NewIntParam(3)},
				},
				map[string]int{"label": 3},
			},
			Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewIntParam(3)},
					{5, NewIntParam(3)},
					{7, NewIntParam(3)},
				},
				map[string]int{},
			},
			0,
		},
		// multiple jump labels
		{
			&Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewStringParam("label")},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{1, NewStringParam("abc")},
					{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0},
			},
			Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewIntParam(0)},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{1, NewIntParam(0)},
					{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			0,
		},
		// no jump label
		{
			&Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			0,
		},
		// unused jump label
		{
			&Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewStringParam("label")},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{1, NewStringParam("abc")},
					{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0, "abcdario": 11},
			},
			Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewIntParam(0)},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{1, NewIntParam(0)},
					{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			0,
		},
		// Fail: jump label that does not exist
		{
			&Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewStringParam("luca")},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{1, NewStringParam("abc")},
					{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0, "abcdario": 11},
			},
			Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewStringParam("luca")},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{1, NewIntParam(0)},
					{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0, "abcdario": 11},
			},
			1,
		},
		// Fail: multiple jump labels that do not exist
		{
			&Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewStringParam("luca")},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{1, NewStringParam("abc")},
					{11, NewIntParam(15)},
				},
				map[string]int{"label": 0, "abcdario": 11},
			},
			Program{
				[]Instruction{
					{3, NewIntParam(1)},
					{1, NewStringParam("luca")},
					{5, NewIntParam(3)},
					{2, NewIntParam(0)},
					{1, NewStringParam("abc")},
					{11, NewIntParam(15)},
				},
				map[string]int{"label": 0, "abcdario": 11},
			},
			2,
		},
	}

	for i, test := range tests {
		errs := test.programBefore.ReplaceLabelsWithNumbers()

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

	program2 := Program{[]Instruction{mockInstruction(), mockInstruction()}, map[string]int{}}
	program2.AddInstruction(mockInstruction())
	if program2.LenInstructions() != 3 {
		t.Errorf("Wrong 1")
	}
}
func mockInstruction() Instruction {
	cmd, _ := NewInstruction(0, NewIntParam(1))
	return *cmd
}
