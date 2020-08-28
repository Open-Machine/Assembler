package instruction

import (
	"strings"

	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/myerrors"
	"github.com/open-machine/assembler/utils"
)

func AssembleInstruction(line string) (*data.Instruction, *myerrors.CustomError) {
	arrayWords := strings.Split(line, " ")
	instructionName := arrayWords[0]

	instructionConfig, exists := instructions[instructionName]
	if !exists {
		err := myerrors.InstructionDoesNotExistError(instructionName)
		return nil, myerrors.NewCodeError(err)
	}

	param, paramErr := instructionConfig.getParam(instructionName, arrayWords)
	if paramErr != nil {
		return nil, paramErr
	}

	instructionPointer, customErr := data.NewInstruction(instructionConfig.code, *param)
	return instructionPointer, customErr
}

func getParamNoParam(instructionName string, words []string) (*data.InstructionParameter, *myerrors.CustomError) {
	if len(words) != 1 {
		remainingParams := getInstructionParams(words)
		err := myerrors.WrongNumberOfParamsError(instructionName, 0, len(remainingParams), remainingParams)
		return nil, myerrors.NewCodeError(err)
	}

	param := data.NewIntParam(0)
	return &param, nil
}

func getSecondWordAsInt(instructionName string, words []string) (*data.InstructionParameter, *myerrors.CustomError) {
	return getSecondWord(instructionName, words, false)
}

func getSecondWordAsIntOrString(instructionName string, words []string) (*data.InstructionParameter, *myerrors.CustomError) {
	return getSecondWord(instructionName, words, true)
}

func getSecondWord(instructionName string, words []string, acceptStringParam bool) (*data.InstructionParameter, *myerrors.CustomError) {
	if len(words) != 2 {
		if len(words) < 2 {
			err := myerrors.WrongNumberOfParamsError(instructionName, 1, 0, []string{})
			return nil, myerrors.NewCodeError(err)
		}

		remainingParams := getInstructionParams(words)
		err := myerrors.WrongNumberOfParamsError(instructionName, 1, len(remainingParams), remainingParams)
		return nil, myerrors.NewCodeError(err)
	}

	strParam := words[1]

	if acceptStringParam && utils.IsValidVarName(strParam) {
		param := data.NewStringParam(strParam)
		return &param, nil
	}

	num, err := utils.StrToPositiveInt(strParam)

	if err != nil {
		var customMsgErr error
		if acceptStringParam {
			customMsgErr = myerrors.InvalidParamLabelOrInt(strParam, err)
		} else {
			customMsgErr = myerrors.InvalidParamInt(strParam, err)
		}
		return nil, myerrors.NewCodeError(customMsgErr)
	}

	param := data.NewIntParam(num)
	return &param, nil
}

func getInstructionParams(words []string) []string {
	return words[1:]
}
