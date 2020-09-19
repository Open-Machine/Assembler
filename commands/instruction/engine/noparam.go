package engine

import (
	"strings"

	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/data"
)

const DEFAULT_INT_PARAM = 0

type NoParamInstruct struct {
	code int
}

func NewNoParamInstruct(c int) NoParamInstruct {
	return NoParamInstruct{code: c}
}

func (assembler NoParamInstruct) Assemble(instructionName string, line string, program *data.Program) *myerrors.CustomError {
	words := strings.Split(line, " ")

	errAmntWords := checkAmntWords(0, instructionName, words)
	if errAmntWords != nil {
		return errAmntWords
	}

	return addInstructionIntParam(assembler.code, DEFAULT_INT_PARAM, program)
}
