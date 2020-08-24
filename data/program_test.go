package data

import (
	"reflect"
	"testing"
)

func TestAddCommand(t *testing.T) {
	program := NewProgram(5)

	if len(program.commands) != 0 {
		t.Errorf("Expected length 0, got: %d", len(program.commands))
	}

	program.AddCommand(Command{0, NewIntParam(0)})

	if len(program.commands) != 1 {
		t.Errorf("Expected length 1, got: %d", len(program.commands))
	}
}

func TestToExecuterSuccess(t *testing.T) {
	program := NewProgram(3)
	program.AddCommand(Command{1, NewIntParam(2)})
	program.AddCommand(Command{15, NewIntParam(7)})
	program.AddCommand(Command{0, NewIntParam(0)})

	got, errors := program.ToExecuter()
	expected := "01020f070000"

	if !(len(errors) == 0 && got == expected) {
		t.Errorf("Expected: '%s', got: '%s'", expected, got)
	}
}

func TestToExecuterFail(t *testing.T) {
	program := NewProgram(3)
	program.AddCommand(Command{1, NewIntParam(2)})
	program.AddCommand(Command{1200, NewIntParam(7)})
	program.AddCommand(Command{0, NewIntParam(0)})

	execCode, errors := program.ToExecuter()

	if len(errors) != 1 {
		t.Errorf("Should result in error because of overflow. Executer code: %s // Errors: %v", execCode, errors)
	}
}

func TestAddGotoLabel(t *testing.T) {
	var tests = []struct {
		program      Program
		expectsError bool
	}{
		{Program{[]Command{}, map[string]int{"abc": 1, "luca": 2}}, false},
		{Program{[]Command{}, map[string]int{"abc": 1, "label": 2}}, true},
		{Program{[]Command{}, map[string]int{"label": 2}}, true},
		{Program{[]Command{}, map[string]int{"a": 2}}, false},
	}

	for i, test := range tests {
		err := test.program.AddGotoLabel("label", 1)
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
		// single goto label
		{
			&Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewStringParam("label")},
					Command{5, NewIntParam(3)},
					Command{7, NewIntParam(3)},
				},
				map[string]int{"label": 3},
			},
			Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewIntParam(3)},
					Command{5, NewIntParam(3)},
					Command{7, NewIntParam(3)},
				},
				map[string]int{},
			},
			0,
		},
		// multiple goto labels
		{
			&Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewStringParam("label")},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{1, NewStringParam("abc")},
					Command{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0},
			},
			Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewIntParam(0)},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{1, NewIntParam(0)},
					Command{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			0,
		},
		// no goto label
		{
			&Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			0,
		},
		// unused goto label
		{
			&Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewStringParam("label")},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{1, NewStringParam("abc")},
					Command{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0, "abcdario": 11},
			},
			Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewIntParam(0)},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{1, NewIntParam(0)},
					Command{11, NewIntParam(15)},
				},
				map[string]int{},
			},
			0,
		},
		// Fail: goto label that does not exist
		{
			&Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewStringParam("luca")},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{1, NewStringParam("abc")},
					Command{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0, "abcdario": 11},
			},
			Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewStringParam("luca")},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{1, NewIntParam(0)},
					Command{11, NewIntParam(15)},
				},
				map[string]int{"abc": 0, "label": 0, "abcdario": 11},
			},
			1,
		},
		// Fail: multiple goto labels that do not exist
		{
			&Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewStringParam("luca")},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{1, NewStringParam("abc")},
					Command{11, NewIntParam(15)},
				},
				map[string]int{"label": 0, "abcdario": 11},
			},
			Program{
				[]Command{
					Command{3, NewIntParam(1)},
					Command{1, NewStringParam("luca")},
					Command{5, NewIntParam(3)},
					Command{2, NewIntParam(0)},
					Command{1, NewStringParam("abc")},
					Command{11, NewIntParam(15)},
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

func TestLenCommands(t *testing.T) {
	program1 := NewProgram(5)
	if program1.LenCommands() != 0 {
		t.Errorf("Wrong 1")
	}

	program2 := Program{[]Command{mockCommand(), mockCommand()}, map[string]int{}}
	program2.AddCommand(mockCommand())
	if program2.LenCommands() != 3 {
		t.Errorf("Wrong 1")
	}
}
func mockCommand() Command {
	cmd, _ := NewCommand(0, NewIntParam(1))
	return *cmd
}
