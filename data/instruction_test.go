package data

import (
	"github.com/open-machine/assembler/config"
	"testing"
)

func TestNewInstructionOverflowValidation(t *testing.T) {
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
		_, err := NewInstruction(test.code, NewIntParam(test.param))
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}
	}
}

func TestNewInstructionWrongStringParam(t *testing.T) {
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
		_, err := NewInstruction(test.code, NewStringParam(test.param))
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}
	}
}

func TestToExecuter(t *testing.T) {
	var tests = []struct {
		instruction      Instruction
		expected     string
		expectsError bool
	}{
		{Instruction{0, NewIntParam(0)}, "0000", false},
		{Instruction{11, NewIntParam(5)}, "0b05", false},
		{Instruction{300, NewIntParam(5)}, "", true},
		{Instruction{5, NewIntParam(300)}, "", true},
		{Instruction{5, NewStringParam("abc")}, "", true},
	}

	for i, test := range tests {
		got, err := test.instruction.toExecuter()
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}

		if test.expected != got {
			t.Errorf("[%d] Expected: %s, Got: %s", i, test.expected, got)
		}
	}
}

func TestNewInstructionTest(t *testing.T) {
	oldTesting := config.Testing
	defer func() { config.Testing = oldTesting }()

	code := 300
	param := NewIntParam(300)

	_, err := NewInstruction(code, param)
	if err == nil {
		t.Errorf("Expected error! NewInstruction should verify params and these params should be wrong to validate the NewInstructionTest function")
	}

	config.Testing = false
	ptrInstructionNil := NewInstructionTest(code, param)
	if ptrInstructionNil != nil {
		t.Errorf("Expected nil instruction, got not nil instruction")
	}

	config.Testing = true
	ptrInstructionNotNil := NewInstructionTest(code, param)
	if ptrInstructionNotNil == nil {
		t.Errorf("Expected nil instruction, got not nil instruction")
	}
}
