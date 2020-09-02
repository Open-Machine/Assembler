package core

import (
	"bufio"
	"fmt"
	"io"

	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/utils"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/helper"
)

func programFromFile(file utils.MyFileInterface) *data.Program {
	reader := bufio.NewReader(file.Reader())
	lineIndex := 1
	program := data.NewProgram(0)

	successful := true

	for {
		line, errRead := reader.ReadString('\n')

		if errRead != nil && errRead != io.EOF {
			helper.LogOtherError(fmt.Sprintf("Error while reading file. Error: %s", errRead.Error()))
			return nil
		}

		jumpLabel, instructionPointer, errAssemble := assembleEntireLine(line)

		if jumpLabel != nil {
			errJumpLabel := program.AddJumpLabel(*jumpLabel, program.LenInstructions())
			if errJumpLabel != nil {
				helper.LogErrorInLine(*myerrors.NewCodeError(errJumpLabel), lineIndex, line)
				return nil
			}
		}

		if errAssemble != nil {
			successful = false
			helper.LogErrorInLine(*errAssemble, lineIndex, line)
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

	if len(errs) > 0 {
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
