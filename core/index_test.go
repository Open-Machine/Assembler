package core

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/open-machine/assembler/utils"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/helper"
)

type testInfoAssembleFile struct {
	//input
	lines        []string
	fileName     string
	execFileName *string
	openSrcErr   error
	openExecErr  error
	// output
	statusReturn      int
	machineCodeOutput *string
	strReturn         string
	substringErrors   []string
}

func TestAssembleFileAux(t *testing.T) {
	config.Out = new(bytes.Buffer)

	tests := []testInfoAssembleFile{
		// Wrong file extension
		getFailTestInfo(
			[]string{""}, "file.java", nil, []string{"extension"}, nil, nil,
		),
		// Error openning the source file
		getFailTestInfo(
			[]string{""}, "file.asm", nil, []string{"open", "file.asm"}, errors.New("Error"), nil,
		),
		// Error on programFromFile
		getFailTestInfo(
			[]string{
				"coding now",
			}, "file.asm", nil, []string{"coding", "exist"}, nil, nil,
		),
		// Error on ReplaceLabelsWithNumbers
		getFailTestInfo(
			[]string{
				"store 0x0",
				"name:",
				"jmp unknown",
			}, "file.asm", nil, []string{"unknown", "exist"}, nil, nil,
		),
		// Not changing file name
		getSuccessTestInfo(
			[]string{
				"kill",
			},
			"file.asm", nil, helper.StringPointer("0900"), "file.run",
		),
		// Changing file name
		getSuccessTestInfo(
			[]string{
				"copy 0x0 # do this",
			},
			"file.asm", helper.StringPointer("name"), helper.StringPointer("0100"), "name.run",
		),
		// Error openning the writer file
		getFailTestInfo(
			[]string{""}, "file.asm", helper.StringPointer("otherfile"), []string{"otherfile"}, nil, errors.New("Error"),
		),
		// Basic valid code with all functionalities
		getSuccessTestInfo(
			[]string{
				"copy 0x0 # do this",
				" name:",
				"store 0xA",
				"kill",
			},
			"file.asm", nil, helper.StringPointer("0100 020a 0900"), "file.run",
		),
	}

	for i, test := range tests {
		config.Err = new(bytes.Buffer)

		str := ""
		for i, line := range test.lines {
			str += line
			if i != len(test.lines)-1 {
				str += "\n"
			}
		}

		fileInput := utils.NewMyBufferAsFile(strings.NewReader(str), test.fileName)
		fileOutput := new(bytes.Buffer)

		var pathNameWritingTo string

		ioReaderFromPath := func(string) (utils.MyFileInterface, error) {
			return &fileInput, test.openSrcErr
		}
		ioWriterFromPath := func(p string) (io.Writer, error) {
			pathNameWritingTo = p
			return fileOutput, test.openExecErr
		}

		statusGot, strGot := AssembleFileAux(test.fileName, test.execFileName, ioReaderFromPath, ioWriterFromPath)

		// Check returns
		if statusGot != test.statusReturn {
			t.Errorf("[%d] Status differs. Got: %d, Expected: %d", i, statusGot, test.statusReturn)
		}
		if strGot != test.strReturn {
			t.Errorf("[%d] Strings differs. Got: '%s', Expected: '%s'", i, strGot, test.strReturn)
		}
		if strGot != "" {
			if pathNameWritingTo != test.strReturn {
				t.Errorf("[%d] Assembler is not actually writing to the right file. Got: '%s', Expected: '%s'", i, pathNameWritingTo, test.strReturn)
			}
		}

		// Check outputs

		stderrStr := config.Err.(*bytes.Buffer).String()

		indexError := strings.Index(stderrStr, "[ERROR]")
		gotError := indexError >= 0

		if test.statusReturn == config.FailStatus && !gotError {
			t.Errorf("[%d] Expected Fail but didn't get error on the stderr. StdErr: %s", i, stderrStr)
		}
		if test.statusReturn == config.SuccessStatus && gotError {
			t.Errorf("[%d] Expected Sucess but did get error on the stderr. StdErr: %s", i, stderrStr)
		}

		if len(test.substringErrors) > 0 {
			errorsStr := stderrStr[indexError:]

			containsAll := true
			for _, word := range test.substringErrors {
				if !strings.Contains(errorsStr, word) {
					containsAll = false
				}
			}
			if !containsAll {
				t.Errorf("[%d] Expected error to have the substrings %v. Actual error: %s", i, test.substringErrors, errorsStr)
			}

			outputStr := fileOutput.String()
			if test.machineCodeOutput != nil && *test.machineCodeOutput != outputStr {
				t.Errorf("[%d] Different machine code expected.", i)
				t.Errorf("\t\tExpected: %s", *test.machineCodeOutput)
				t.Errorf("\t\tGot: %s", outputStr)
			}
		}
	}
}
func getSuccessTestInfo(lines []string, fileName string, execFileName *string, machineCodeOutput *string, strReturn string) testInfoAssembleFile {
	return testInfoAssembleFile{
		lines:        lines,
		openSrcErr:   nil,
		openExecErr:  nil,
		fileName:     fileName,
		execFileName: execFileName,

		statusReturn:      config.SuccessStatus,
		machineCodeOutput: machineCodeOutput,
		strReturn:         strReturn,
		substringErrors:   nil,
	}
}
func getFailTestInfo(lines []string, fileName string, execFileName *string, substringErrors []string, openSrcErr error, openExecErr error) testInfoAssembleFile {
	return testInfoAssembleFile{
		lines:        lines,
		openSrcErr:   openSrcErr,
		openExecErr:  openExecErr,
		fileName:     fileName,
		execFileName: execFileName,

		statusReturn:      config.FailStatus,
		machineCodeOutput: nil,
		strReturn:         "",
		substringErrors:   substringErrors,
	}
}
