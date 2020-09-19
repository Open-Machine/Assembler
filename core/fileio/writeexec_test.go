package fileio

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/data"
)

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
			data.ProgramFromInstructionsAndLabels([]data.Instruction{
				*newInstruction(0x2, 1),
				*newInstruction(0x2, 12),
			}, map[string]int{}),
			"v2.0 raw\n00002001200c",
			0,
			false,
		},
		{
			data.ProgramFromInstructionsAndLabels([]data.Instruction{
				*newInstruction(0x2, 1),
				*newInstruction(0x2, 12),
				*newInstruction(0xD, 15),
			}, map[string]int{}),
			"v2.0 raw\n00002001200cd00f",
			0,
			false,
		},
		{
			data.ProgramFromInstructionsAndLabels([]data.Instruction{
				*newInstruction(0x2, 1),
				*newInstruction(0x222, 12),
				*newInstruction(0xD, 115),
			}, map[string]int{}),
			"",
			1,
			true,
		},
		{
			data.ProgramFromInstructionsAndLabels([]data.Instruction{
				*newInstruction(0x2, 1),
				*newInstruction(0x2, 1),
				*newInstructionStr(0x2, "a"),
			}, map[string]int{}),
			"",
			1,
			true,
		},
	}

	for i, test := range tests {
		fileWriter := new(bytes.Buffer)
		got := WriteExecProgram(test.param, "File", fileWriter)

		if !reflect.DeepEqual(test.expectedCode, got) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedCode, got)
		}

		gotFileStr := fileWriter.String()
		if gotFileStr != test.expectedFileStr {
			t.Errorf("[%d] Expected file str: %v, Got file str: %v", i, test.expectedFileStr, gotFileStr)
		}

		stderrStr := config.Err.(*bytes.Buffer).String()
		gotErr := stderrStr != ""
		if test.expectsErr != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t // ", i, test.expectsErr, gotErr)
			t.Errorf("\t\t StdErr: %s", stderrStr)
		}
	}
}
func newInstruction(cmdCode int, intParam int) *data.Instruction {
	param, _ := data.NewIntParam(intParam)
	got, _ := data.NewInstruction(cmdCode, *param)
	return got
}
func newInstructionStr(cmdCode int, strParam string) *data.Instruction {
	param, _ := data.NewStringParam(strParam)
	got, _ := data.NewInstruction(cmdCode, *param)
	return got
}
