package engine

import (
	"strings"

	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/utils"
)

const ACCEPT_STRING_PARAM = true

type OneParamInstruct struct {
	code              int
	acceptStringParam bool
}

func NewOneParamInstruct(code int, acceptStringParam bool) OneParamInstruct {
	return OneParamInstruct{code: code, acceptStringParam: acceptStringParam}
}

func (assembler OneParamInstruct) Assemble(instructionName string, line string, program *data.Program) *myerrors.CustomError {
	words := strings.Split(line, " ")

	errAmntWords := checkAmntWords(1, instructionName, words)
	if errAmntWords != nil {
		return errAmntWords
	}

	strParam := words[1]

	if assembler.acceptStringParam && utils.IsValidName(strParam) {
		return addInstructionStrParam(assembler.code, strParam, program)
	}

	num, err := utils.StrToPositiveInt(strParam)

	if err != nil {
		var customMsgErr error
		if assembler.acceptStringParam {
			customMsgErr = myerrors.InvalidParamLabelOrInt(strParam, err)
		} else {
			customMsgErr = myerrors.InvalidParamInt(strParam, err)
		}
		return myerrors.NewCodeError(customMsgErr)
	}

	return addInstructionIntParam(assembler.code, num, program)
}
