package cli

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/core"
	"github.com/open-machine/assembler/utils"
)

// TODO: test if the shell commands are printing right or returning the right AssembleInstruction (change os.Args)

func TestAssemblyCompiledExample(t *testing.T) {
	config.Err = new(bytes.Buffer)
	fileOutput := new(bytes.Buffer)

	assemblyLines := GetAssemblyExample()
	assemblyCodeStr := ""
	for _, line := range assemblyLines {
		assemblyCodeStr += line + "\n"
	}

	exampleFileName := "file.asm"

	ioReaderFromPath := func(string) (utils.MyFileInterface, error) {
		reader := strings.NewReader(assemblyCodeStr)
		myFile := utils.NewMyBufferAsFile(reader, exampleFileName)
		return &myFile, nil
	}
	ioWriterFromPath := func(string) (io.Writer, error) {
		return fileOutput, nil
	}

	statusGot, strGot := core.AssembleFileAux("file.asm", nil, ioReaderFromPath, ioWriterFromPath)

	if statusGot != config.SuccessStatus {
		t.Errorf("Expected Success Status, but got %d", statusGot)
	}

	fileOutputStr := fileOutput.String()
	replacedFileOutputStr := strings.ReplaceAll(fileOutputStr, " ", "")
	if replacedFileOutputStr == "" {
		t.Errorf("File shouldnt be empty. File output str: '%s'", fileOutputStr)
	}

	expectFileName := exampleFileName + config.AssemblyFileExtension
	if expectFileName == strGot {
		t.Errorf("Expected same file name with '.run' extension. File name: '%s', Got: '%s'", expectFileName, strGot)
	}

	stderrStr := config.Err.(*bytes.Buffer).String()
	if stderrStr != "" {
		t.Errorf("No errors expected, but stderr is not empty: '%s'", stderrStr)
	}
}
