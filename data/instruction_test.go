package data

import (
	"testing"

	"github.com/open-machine/assembler/config"
)

func TestNewInstructionOverflowValidation(t *testing.T) {
	var tests = []struct {
		code         int
		param        int
		expectsError bool
	}{
		{0, 0, false},
		// neg
		{-1, 0, true},
		{0, -1, true},
		// large num
		{0, 5000, true},
		{1000, 0, true},
		// line between right and wrong
		{15, 0, false},
		{16, 0, true},
		{0, 4095, false},
		{0, 4096, true},
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
		instruction  Instruction
		expected     string
		expectsError bool
	}{
		{Instruction{0, NewIntParam(0)}, "0000", false},
		{Instruction{11, NewIntParam(5)}, "b005", false},
		{Instruction{5, NewIntParam(300)}, "512c", false},
		{Instruction{5000, NewIntParam(5)}, "", true},
		{Instruction{5, NewIntParam(5000)}, "", true},
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
