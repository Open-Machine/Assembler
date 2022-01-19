package core

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/open-machine/assembler/utils"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/helper"
)

func TestProgramFromFile(t *testing.T) {
	var tests = []struct {
		lines      []string
		expected   *data.Program
		expectsErr bool
	}{
		// blank lines
		{
			[]string{
				"       ",
				" ",
				"    ",
			},
			newProgramPointer([]data.Instruction{}, map[string]int{}, map[string]data.Variable{}),
			false,
		},
		// variables and code
		{
			[]string{
				"@VAR",
				"var2 = 3",
				"variable = 1",
				"",
				"@CODE",
				"store variable",
				"for:",
				"store var2",
				"end:",
			},
			newProgramPointer(
				[]data.Instruction{
					*data.NewInstructionTest(0x2, "variable", 0, data.VariableParam),
					*data.NewInstructionTest(0x2, "var2", 0, data.VariableParam),
				},
				map[string]int{
					"for": 1,
					"end": 2,
				},
				map[string]data.Variable{
					"var2":     createVariable("var2", 0, 3),
					"variable": createVariable("variable", 1, 1),
				},
			),
			false,
		},
		// only method
		{
			[]string{
				"@CODE",
				"store var",
			},
			newProgramPointer(
				[]data.Instruction{
					*data.NewInstructionTest(0x2, "var", 0, data.VariableParam),
				},
				map[string]int{},
				map[string]data.Variable{},
			),
			false,
		},
		// wrong method
		{
			[]string{
				"@CODE",
				"mov 1",
			},
			nil,
			true,
		},
		// invalid status
		{
			[]string{
				"@CODE",
				"store variable",
				"@VAR",
			},
			nil,
			true,
		},
	}

	for i, test := range tests {
		str := ""
		for i, line := range test.lines {
			str += line
			if i != len(test.lines)-1 {
				str += "\n"
			}
		}

		config.Out = new(bytes.Buffer)
		config.Err = new(bytes.Buffer)

		f := utils.NewMyBufferAsFile(strings.NewReader(str), "file.asm")
		got := programFromFile(&f)

		if !helper.SafeIsEqualProgramPointer(test.expected, got) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expected, got)
		}

		stderrStr := config.Err.(*bytes.Buffer).String()
		gotErr := stderrStr != ""
		if test.expectsErr != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotErr)
			t.Errorf("\t\t StdErr: %s", stderrStr)
		}
	}
}
func newProgramPointer(instructions []data.Instruction, jumpLabelsDict map[string]int, variablesDict map[string]data.Variable) *data.Program {
	prog := data.NewCompleteProgram(instructions, jumpLabelsDict, variablesDict)
	return &prog
}
func createVariable(name string, index int, initialValue uint) data.Variable {
	variable, _ := data.NewVariable(name, index, initialValue)
	return *variable
}

func TestWriteAssembledFile(t *testing.T) {
	config.Out = new(bytes.Buffer)
	config.Err = new(bytes.Buffer)

	config.Testing = true

	var tests = []struct {
		param           data.Program
		expectedFileStr string
		expectedCode    int
		expectsErr      bool
	}{
		{
			data.NewCompleteProgram(
				[]data.Instruction{
					*newInstruction(0x2, 1),
					*newInstruction(0x2, 12),
				}, map[string]int{}, map[string]data.Variable{},
			),
			"v2.0 raw\n0000 2001 200c 4093*0 ",
			0,
			false,
		},
		{
			data.NewCompleteProgram(
				[]data.Instruction{
					*newInstruction(0x2, 1),
					*newInstruction(0x2, 12),
					*newInstruction(0xD, 15),
				}, map[string]int{}, map[string]data.Variable{},
			),
			"v2.0 raw\n0000 2001 200c d00f 4092*0 ",
			0,
			false,
		},
		{
			data.NewCompleteProgram(
				[]data.Instruction{
					*newInstruction(0x2, 1),
					*newInstruction(0x222, 12),
					*newInstruction(0xD, 115),
				}, map[string]int{}, map[string]data.Variable{},
			),
			"",
			1,
			true,
		},
	}

	for i, test := range tests {
		fileWriter := new(bytes.Buffer)
		got := writeExecProgram(test.param, "File", fileWriter)

		if !reflect.DeepEqual(test.expectedCode, got) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedCode, got)
		}

		gotFileStr := fileWriter.String()
		if gotFileStr != test.expectedFileStr {
			t.Errorf("[%d] Expected file str: '%s', Got file str: '%s'", i, test.expectedFileStr, gotFileStr)
		}

		stderrStr := config.Err.(*bytes.Buffer).String()
		gotErr := stderrStr != ""
		if test.expectsErr != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t // StdErr: %s", i, test.expectsErr, gotErr, stderrStr)
		}
	}
}
func newInstruction(cmdCode int, param int) *data.Instruction {
	return data.NewInstructionTest(cmdCode, "", param, 0)
}
