package core

import (
	"assembler/data"
	"assembler/utils/helper"
	"reflect"
	"strings"
	"testing"
)

func TestGetNoParam(t *testing.T) {
	got, err := getParamNoParam("nop", []string{"nop"})
	if !(err == nil && !got.IsStr && got.Num == 0) {
		t.Errorf("Wrong")
	}
}

func TestGetSecondParamAsInt(t *testing.T) {
	var tests = []struct {
		line       string
		expected   *data.CommandParameter
		expectsErr bool
	}{
		// Decimal Number
		{"mov 1", newCmdIntParam(1), false},
		// Hexadecimal Number
		{"mov 0x1a", newCmdIntParam(26), false},
		{"mov 0x001", newCmdIntParam(1), false},
		{"mov 0x0f", newCmdIntParam(15), false},
		{"mov 0xff", newCmdIntParam(255), false},
		{"mov 0x0ff", newCmdIntParam(255), false},
		// Variable
		{"mov 0xx0ff", nil, true},
		{"mov x1", nil, true},
		{"mov 0x1g", nil, true},
		{"mov 1a", nil, true},
		// Words
		{"mov", nil, true},
		{"mov a b", nil, true},
		{"mov a b c", nil, true},
	}

	for i, test := range tests {
		arrayWords := strings.Split(test.line, " ")
		got, err := getSecondWord(arrayWords[0], arrayWords, false)
		gotError := err != nil

		if test.expectsErr != gotError {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotError)
		}

		if !helper.SafeIsEqualCommandParamPointer(test.expected, got) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expected, got)
		}

		if got != nil && got.IsStr {
			t.Errorf("[%d] Expecting int parameter", i)
		}
	}
}

func TestGetSecondParamAsIntOrString(t *testing.T) {
	var tests = []struct {
		line       string
		expected   *data.CommandParameter
		expectsErr bool
	}{
		// Decimal Number
		{"jmp 1", newCmdIntParam(1), false},
		// Hexadecimal Number
		{"jmp 0x001", newCmdIntParam(1), false},
		{"jmp 0x0f", newCmdIntParam(15), false},
		{"jmp 0xff", newCmdIntParam(255), false},
		{"jmp 0x0ff", newCmdIntParam(255), false},
		// Variable
		{"jmp a8", newCmdStringParam("a8"), false},
		{"jmp x1", newCmdStringParam("x1"), false},
		{"jmp a1", newCmdStringParam("a1"), false},
		// Errors 1 param
		{"jmp 0xx0ff", nil, true},
		{"jmp 0x1g", nil, true},
		{"jmp 1a", nil, true},
		// Erros amnt params
		{"jmp", nil, true},
		{"jmp a b", nil, true},
		{"jmp a b c", nil, true},
	}

	for i, test := range tests {
		arrayWords := strings.Split(test.line, " ")
		got, err := getSecondWord(arrayWords[0], arrayWords, true)
		gotError := err != nil

		if test.expectsErr != gotError {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotError)
		}

		if !helper.SafeIsEqualCommandParamPointer(test.expected, got) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expected, got)
		}

		if got != nil && test.expected != nil && got.IsStr != test.expected.IsStr {
			t.Errorf("[%d] Expected IsStr: %t, Got IsStr: %t", i, test.expected.IsStr, got.IsStr)
		}
	}
}
func newCmdIntParam(num int) *data.CommandParameter {
	param := data.NewIntParam(num)
	return &param
}
func newCmdStringParam(str string) *data.CommandParameter {
	param := data.NewStringParam(str)
	return &param
}

func TestAssembleCommand(t *testing.T) {
	if len(commands) != 12 {
		t.Errorf("Tests were not updated")
	}

	var tests = []struct {
		line       string
		expected   *data.Command
		expectsErr bool
	}{
		// Success Number
		{"nop", getCommand(0x0, 0), false},
		{"copy 0x10", getCommand(0x1, 16), false},
		{"store 0x10", getCommand(0x2, 16), false},
		{"add 10", getCommand(0x3, 10), false},
		{"sub 10", getCommand(0x4, 10), false},
		{"input 7", getCommand(0x7, 7), false},
		{"output 8", getCommand(0x8, 8), false},
		{"kill", getCommand(0x9, 0), false},
		{"jmp 0x8", getCommand(0xA, 8), false},
		{"jg 0x8", getCommand(0xB, 8), false},
		{"je 0x8", getCommand(0xD, 8), false},
		{"jl 0x8", getCommand(0xF, 8), false},
		// Success Label
		{"jmp label", getCommandStr(0xA, "label"), false},
		{"jg label", getCommandStr(0xB, "label"), false},
		{"je label", getCommandStr(0xD, "label"), false},
		{"jl label", getCommandStr(0xF, "label"), false},
		// Fail: Wrong Command
		{"nope", nil, true},
		// Fail: No label as param
		{"copy label", nil, true},
		{"store label", nil, true},
		{"add label", nil, true},
		{"sub label", nil, true},
		{"input label", nil, true},
		{"output label", nil, true},
		// Fail: Wrong param
		{"add 1x10", nil, true},
		// Fail: Amnt params
		{"kill 0", nil, true},
		{"output", nil, true},
		{"output 8 1", nil, true},
	}

	for i, test := range tests {
		got, err := AssembleCommand(test.line)
		gotError := err != nil

		if test.expectsErr != gotError {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotError)
		}
		if !helper.SafeIsEqualCommandPointer(test.expected, got) {
			t.Errorf("Command expected is: %v, Got expected is: %v", test.expected, got)
		}
	}
}
func getCommand(code int, param int) *data.Command {
	cmd, _ := data.NewCommand(code, data.NewIntParam(param))
	return cmd
}
func getCommandStr(code int, param string) *data.Command {
	cmd, _ := data.NewCommand(code, data.NewStringParam(param))
	return cmd
}

func TestGetCommandParams(t *testing.T) {
	expected := []string{"0x1", "1", "label"}
	got := getCommandParams([]string{"mov", "0x1", "1", "label"})
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error")
	}
}
