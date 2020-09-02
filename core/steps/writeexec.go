package steps

import (
	"bufio"
	"fmt"
	"io"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/helper"
)

func WriteExecProgram(program data.Program, execFileName string, execFile io.Writer) int {
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
