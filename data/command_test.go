package data

import (
	"assembler/config"
	"testing"
)

func TestNewCommandOverflowValidation(t *testing.T) {
	var tests = []struct {
		code         int
		param        int
		expectsError bool
	}{
		{0, 0, false},
		{-1, 0, true},
		{0, -1, true},
		{1000, 1000, true},
		{255, 0, false},
		{256, 0, true},
		{0, 255, false},
		{0, 256, true},
	}

	for i, test := range tests {
		_, err := NewCommand(test.code, NewIntParam(test.param))
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}
	}
}

func TestNewCommandWrongStringParam(t *testing.T) {
	var tests = []struct {
		code         int
		param        string
		expectsError bool
	}{
		{0, "", true},
		{0, "1a", true},
		{0, "a1", false},
	}

	for i, test := range tests {
		_, err := NewCommand(test.code, NewStringParam(test.param))
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}
	}
}

func TestToExecuter(t *testing.T) {
	var tests = []struct {
		command      Command
		expected     string
		expectsError bool
	}{
		{Command{0, NewIntParam(0)}, "0000", false},
		{Command{11, NewIntParam(5)}, "0b05", false},
		{Command{300, NewIntParam(5)}, "", true},
		{Command{5, NewIntParam(300)}, "", true},
		{Command{5, NewStringParam("abc")}, "", true},
	}

	for i, test := range tests {
		got, err := test.command.toExecuter()
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}

		if test.expected != got {
			t.Errorf("[%d] Expected: %s, Got: %s", i, test.expected, got)
		}
	}
}

func TestNewCommandTest(t *testing.T) {
	oldTesting := config.Testing
	defer func() { config.Testing = oldTesting }()

	code := 300
	param := NewIntParam(300)

	_, err := NewCommand(code, param)
	if err == nil {
		t.Errorf("Expected error! NewCommand should verify params and these params should be wrong to validate the NewCommandTest function")
	}

	config.Testing = false
	ptrCommandNil := NewCommandTest(code, param)
	if ptrCommandNil != nil {
		t.Errorf("Expected nil command, got not nil command")
	}

	config.Testing = true
	ptrCommandNotNil := NewCommandTest(code, param)
	if ptrCommandNotNil == nil {
		t.Errorf("Expected nil command, got not nil command")
	}
}
