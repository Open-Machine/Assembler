package instruction

import (
	"reflect"
	"strings"
	"testing"

	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/helper"
)

func TestGetParamString(t *testing.T) {
	var tests = []struct {
		line       string
		expected   string
		expectsErr bool
	}{
		// Variable
		{"jmp label", "label", false},
		{"store variable", "variable", false},
		// Erros amnt params
		{"jmp", "", true},
		{"jmp a b", "", true},
		{"jmp a b c", "", true},
	}

	for i, test := range tests {
		arrayWords := strings.Split(test.line, " ")
		got, err := getParam(arrayWords[0], arrayWords)
		gotError := err != nil

		if test.expectsErr != gotError {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotError)
		}

		if test.expected != got {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expected, got)
		}
	}
}

func TestAssembleInstruction(t *testing.T) {
	if len(instructions) != 12 {
		t.Errorf("Tests were not updated")
	}

	var tests = []struct {
		line       string
		expected   *data.Instruction
		expectsErr bool
	}{
		// Success Number
		{"nop", getInstruction(0x0, "", "no-param"), false},
		{"copy variable", getInstruction(0x1, "variable", "variable"), false},
		{"store var", getInstruction(0x2, "var", "variable"), false},
		{"add a", getInstruction(0x3, "a", "variable"), false},
		{"sub b", getInstruction(0x4, "b", "variable"), false},
		{"input c", getInstruction(0x7, "c", "variable"), false},
		{"output c", getInstruction(0x8, "c", "variable"), false},
		{"kill", getInstruction(0x9, "", "no-param"), false},
		{"jmp label", getInstruction(0xA, "label", "jump"), false},
		{"jg labeljmp", getInstruction(0xB, "labeljmp", "jump"), false},
		{"je jaa", getInstruction(0xD, "jaa", "jump"), false},
		{"jl l", getInstruction(0xF, "l", "jump"), false},
		// Fail: Wrong Instruction
		{"nope", nil, true},
		// Fail: Cant be a number
		{"add 10", nil, true},
		{"add 0x10", nil, true},
		// Fail: Wrong variable name
		{"jmp f_A", nil, true},
		// Fail: reserved word
		{"add nop", nil, true},
		// Fail: Amnt params
		{"kill 0", nil, true},
		{"output", nil, true},
		{"output 8 1", nil, true},
	}

	for i, test := range tests {
		got, err := AssembleInstruction(test.line)
		gotError := err != nil

		if test.expectsErr != gotError {
			t.Errorf("[%d] Expected error: %t, Got error: %t, Str: '%s'", i, test.expectsErr, gotError, test.line)
		}
		if !helper.SafeIsEqualInstructionPointer(test.expected, got) {
			t.Errorf("Instruction expected is: %v, Got expected is: %v", test.expected, got)
		}
	}
}
func getInstruction(code int, param string, t string) *data.Instruction {
	switch t {
	case "no-param":
		instruction := data.NewInstructionWithoutParam(code)
		return &instruction
	case "variable":
		cmd, _ := data.NewVariableInstruction(code, param)
		return cmd
	case "jump":
		cmd, _ := data.NewJumpInstruction(code, param)
		return cmd
	}
	return nil
}

func TestGetInstructionParams(t *testing.T) {
	expected := []string{"0x1", "1", "label"}
	got := getInstructionParams([]string{"mov", "0x1", "1", "label"})
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error")
	}
}
