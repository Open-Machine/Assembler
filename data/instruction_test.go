package data

import (
	"testing"

	"github.com/open-machine/assembler/config"
)

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
		_, err1 := NewVariableInstruction(test.code, test.param)
		_, err2 := NewJumpInstruction(test.code, test.param)

		gotErr1 := err1 != nil
		if test.expectsError != gotErr1 {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr1)
		}

		gotErr2 := err2 != nil
		if test.expectsError != gotErr2 {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr2)
		}
	}
}

func TestToExecuter(t *testing.T) {
	var tests = []struct {
		instruction  Instruction
		expected     string
		expectsError bool
	}{
		{Instruction{0, "", 0, 0}, "0000", false},
		{Instruction{11, "", 5, 0}, "b005", false},
		{Instruction{5, "", 300, 0}, "512c", false},
		{Instruction{5000, "", 5, 0}, "", true},
		{Instruction{5, "", 5000, 0}, "", true},
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
	config.TestSetup()

	code := 300
	param := 300

	ptrInstructionNotNil := NewInstructionTest(code, "", param, 0)
	if ptrInstructionNotNil == nil {
		t.Errorf("Expected not nil instruction, got nil instruction")
	}

	config.Testing = false
	ptrInstructionNil := NewInstructionTest(code, "", param, 0)
	if ptrInstructionNil != nil {
		t.Errorf("Expected nil instruction, got not nil instruction")
	}

	config.Testing = true
	ptrInstructionNotNil2 := NewInstructionTest(code, "", param, 0)
	if ptrInstructionNotNil2 == nil {
		t.Errorf("Expected not nil instruction, got nil instruction")
	}
}
