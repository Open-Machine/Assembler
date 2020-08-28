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

func TestToExecuterSuccess(t *testing.T) {
	program := NewProgram(3)
	program.AddInstruction(Instruction{1, NewIntParam(2)})
	program.AddInstruction(Instruction{15, NewIntParam(7)})
	program.AddInstruction(Instruction{0, NewIntParam(0)})

	got, errors := program.ToExecuter()
	expected := "01020f070000"

	if !(len(errors) == 0 && got == expected) {
		t.Errorf("Expected: '%s', got: '%s'", expected, got)
	}
}

func TestToExecuterFail(t *testing.T) {
	program := NewProgram(3)
	program.AddInstruction(Instruction{1, NewIntParam(2)})
	program.AddInstruction(Instruction{1200, NewIntParam(7)})
	program.AddInstruction(Instruction{0, NewIntParam(0)})

	execCode, errors := program.ToExecuter()

	if len(errors) != 1 {
		t.Errorf("Should result in error because of overflow. Executer code: %s // Errors: %v", execCode, errors)
	}
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
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewStringParam("label")},
					Instruction{5, NewIntParam(3)},
					Instruction{7, NewIntParam(3)},
				},
				map[string]int{"label": 3},
			},
			Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewIntParam(3)},
					Instruction{5, NewIntParam(3)},
					Instruction{7, NewIntParam(3)},
				},
				map[string]int{},
			},
			0,
		},
		// multiple jump labels
		{
			&Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewStringParam("label")},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{1, NewStringParam("abc")},
					Instruction{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0},
			},
			Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewIntParam(0)},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{1, NewIntParam(0)},
					Instruction{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			0,
		},
		// no jump label
		{
			&Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			0,
		},
		// unused jump label
		{
			&Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewStringParam("label")},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{1, NewStringParam("abc")},
					Instruction{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0, "abcdario": 11},
			},
			Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewIntParam(0)},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{1, NewIntParam(0)},
					Instruction{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			0,
		},
		// Fail: jump label that does not exist
		{
			&Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewStringParam("luca")},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{1, NewStringParam("abc")},
					Instruction{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0, "abcdario": 11},
			},
			Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewStringParam("luca")},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{1, NewIntParam(0)},
					Instruction{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0, "abcdario": 11},
			},
			1,
		},
		// Fail: multiple jump labels that do not exist
		{
			&Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewStringParam("luca")},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{1, NewStringParam("abc")},
					Instruction{11, NewIntParam(15)},
				},
				map[string]int{"label": 0, "abcdario": 11},
			},
			Program{
				[]Instruction{
					Instruction{3, NewIntParam(1)},
					Instruction{1, NewStringParam("luca")},
					Instruction{5, NewIntParam(3)},
					Instruction{2, NewIntParam(0)},
					Instruction{1, NewStringParam("abc")},
					Instruction{11, NewIntParam(15)},
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
