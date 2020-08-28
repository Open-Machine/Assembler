package core

import (
	"bufio"
	"fmt"
	"io"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/utils/helper"
)

func programFromFile(file io.Reader) *data.Program {
	reader := bufio.NewReader(file)
	lineIndex := 1
	program := data.NewProgram(0)

	successful := true

	for {
		line, errRead := reader.ReadString('\n')

		if errRead != nil && errRead != io.EOF {
			helper.LogOtherError(fmt.Sprintf("Error while reading file. Error: %s", errRead.Error()))
			return nil
		}

		jumpLabel, instructionPointer, err := assembleEntireLine(line)

		if jumpLabel != nil {
			program.AddJumpLabel(*jumpLabel, program.LenInstructions())
		}

		if err != nil {
			successful = false
			helper.LogErrorInLine(*err, lineIndex, line)
		} else if instructionPointer != nil {
			program.AddInstruction(*instructionPointer)
		}

		if errRead == io.EOF {
			break
		}
		lineIndex++
	}

	if !successful {
		return nil
	}
	return &program
}

func writeExecProgram(program data.Program, execFileName string, execFile io.Writer) int {
	writer := bufio.NewWriter(execFile)
	defer writer.Flush()

	execStr, errs := program.ToExecuter()

	if errs != nil && len(errs) > 0 {
		for _, err := range errs {
			// TODO: infrastructure to get line
			helper.LogErrorInLine(err, 0, "")
		}
		return config.FailStatus
	}

	_, err := writer.WriteString(execStr)
	if err != nil {
		helper.LogOtherError(fmt.Sprintf("Could not write to file %s \n", execFileName))
		return config.FailStatus
	}

	return config.SuccessStatus
}
