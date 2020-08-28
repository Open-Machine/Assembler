package core

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/utils/helper"
)

func TestWriteAssembledFile(t *testing.T) {
	stdout := helper.Out
	helper.Out = new(bytes.Buffer)
	defer func() { helper.Out = stdout }()

	stderr := helper.Err
	helper.Err = new(bytes.Buffer)
	defer func() { helper.Err = stderr }()

	oldTesting := config.Testing
	config.Testing = true
	defer func() { config.Testing = oldTesting }()

	var tests = []struct {
		param           data.Program
		expectedFileStr string
		expectedCode    int
		expectsErr      bool
	}{
		{
			data.ProgramFromInstructionsAndLabels([]data.Instruction{
				*newInstruction(0x2, 1),
				*newInstruction(0x2, 12),
			}, map[string]int{}),
			"0201020c",
			0,
			false,
		},
		{
			data.ProgramFromInstructionsAndLabels([]data.Instruction{
				*newInstruction(0x2, 1),
				*newInstruction(0x2, 12),
				*newInstruction(0xD, 15),
			}, map[string]int{}),
			"0201020c0d0f",
			0,
			false,
		},
		{
			data.ProgramFromInstructionsAndLabels([]data.Instruction{
				*newInstruction(0x2, 1),
				*data.NewInstructionTest(0x222, data.NewIntParam(12)),
				*newInstruction(0xD, 115),
			}, map[string]int{}),
			"",
			1,
			true,
		},
		{
			data.ProgramFromInstructionsAndLabels([]data.Instruction{
				*newInstruction(0x2, 1),
				*data.NewInstructionTest(0x2, data.NewStringParam("a")),
			}, map[string]int{}),
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
			t.Errorf("[%d] Expected file str: %v, Got file str: %v", i, test.expectedFileStr, gotFileStr)
		}

		stderrStr := helper.Err.(*bytes.Buffer).String()
		gotErr := stderrStr != ""
		if test.expectsErr != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t // ", i, test.expectsErr, gotErr)
			t.Errorf("\t\t StdErr: %s", stderrStr)
		}
	}
}

func newInstruction(cmdCode int, param int) *data.Instruction {
	got, _ := data.NewInstruction(cmdCode, data.NewIntParam(param))
	return got
}

func TestProgramFromFile(t *testing.T) {
	stdout := helper.Out
	defer func() { helper.Out = stdout }()
	stderr := helper.Err
	defer func() { helper.Err = stderr }()

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
			newProgramPointer([]data.Instruction{}, map[string]int{}),
			false,
		},
		// blank lines and right methods
		{
			[]string{
				"",
				"  ",
				"store 1",
				"store 0xC",
			},
			newProgramPointer(
				[]data.Instruction{
					*newInstruction(0x2, 1),
					*newInstruction(0x2, 12),
				},
				map[string]int{},
			),
			false,
		},
		// wrong method
		{
			[]string{
				"mov 1",
			},
			nil,
			true,
		},
		// wrong params
		{
			[]string{
				"store",
			},
			nil,
			true,
		},
		// jump labels always on same line and instructionIndex!=lineIndex (successful)
		{
			[]string{
				"",
				"  ",
				"label1: store 1",
				"store 0xC",
				"label2: copy 0xF",
			},
			newProgramPointer(
				[]data.Instruction{
					*newInstruction(0x2, 1),
					*newInstruction(0x2, 12),
					*newInstruction(0x1, 15),
				},
				map[string]int{
					"label1": 0,
					"label2": 2,
				},
			),
			false,
		},
		// labels always alone on the line (successful)
		{
			[]string{
				"label1 :  ",
				"store 1",
				"store 0xC",
				"jmping:",
				"store 0xC",
			},
			newProgramPointer(
				[]data.Instruction{
					*newInstruction(0x2, 1),
					*newInstruction(0x2, 12),
					*newInstruction(0x2, 12),
				},
				map[string]int{
					"label1": 0,
					"jmping": 2,
				},
			),
			false,
		},
		// jump labels mixed on same line and on the line above (successful)
		{
			[]string{
				"store 1",
				"label1:",
				"store 1",
				"store 0xC",
				"label2: copy 0xF",
				"kill",
			},
			newProgramPointer(
				[]data.Instruction{
					*newInstruction(0x2, 1),
					*newInstruction(0x2, 1),
					*newInstruction(0x2, 12),
					*newInstruction(0x1, 15),
					*newInstruction(0x9, 0),
				},
				map[string]int{
					"label1": 1,
					"label2": 3,
				},
			),
			false,
		},
		// invalid label (fail)
		{
			[]string{
				"1label1 :  ",
				"store 1",
				"store 0xC",
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

		helper.Out = new(bytes.Buffer)
		helper.Err = new(bytes.Buffer)

		got := programFromFile(strings.NewReader(str))

		if !helper.SafeIsEqualProgramPointer(test.expected, got) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expected, got)
		}

		stderrStr := helper.Err.(*bytes.Buffer).String()
		gotErr := stderrStr != ""
		if test.expectsErr != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotErr)
			t.Errorf("\t\t StdErr: %s", stderrStr)
		}
	}
}
func newProgramPointer(instructions []data.Instruction, jumpLabelsDict map[string]int) *data.Program {
	prog := data.ProgramFromInstructionsAndLabels(instructions, jumpLabelsDict)
	return &prog
}
