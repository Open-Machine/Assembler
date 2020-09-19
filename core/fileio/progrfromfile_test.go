package fileio

import (
	"bytes"
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
			nil,
			true,
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

		config.Out = new(bytes.Buffer)
		config.Err = new(bytes.Buffer)

		f := utils.NewMyBufferAsFile(strings.NewReader(str), "file.asm")
		got := ProgramFromFile(&f)

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
func newProgramPointer(instructions []data.Instruction, jumpLabelsDict map[string]int) *data.Program {
	prog := data.ProgramFromInstructionsAndLabels(instructions, jumpLabelsDict)
	return &prog
}
