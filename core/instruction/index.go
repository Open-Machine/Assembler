package instruction

import (
	"strings"

	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/data"
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

	switch instructionConfig.parameterType {
	case data.VariableParam:
		param, paramErr := getParam(instructionName, arrayWords)
		if paramErr != nil {
			return nil, paramErr
		}
		err := utils.CheckVariableName(param)
		if err != nil {
			return nil, myerrors.NewCodeError(err)
		}
		return data.NewVariableInstruction(instructionConfig.code, param)
	case data.JumpLabelParam:
		param, paramErr := getParam(instructionName, arrayWords)
		if paramErr != nil {
			return nil, paramErr
		}
		err := utils.CheckVariableName(param)
		if err != nil {
			return nil, myerrors.NewCodeError(err)
		}
		return data.NewJumpInstruction(instructionConfig.code, param)
	case data.NoParam:
	}
	err := getNoParam(instructionName, arrayWords)
	if err != nil {
		return nil, err
	}
	instruction := data.NewInstructionWithoutParam(instructionConfig.code)
	return &instruction, nil
}

func getParam(instructionName string, words []string) (string, *myerrors.CustomError) {
	if len(words) != 2 {
		if len(words) < 2 {
			err := myerrors.WrongNumberOfParamsError(instructionName, 1, 0, []string{})
			return "", myerrors.NewCodeError(err)
		}

		remainingParams := getInstructionParams(words)
		err := myerrors.WrongNumberOfParamsError(instructionName, 1, len(remainingParams), remainingParams)
		return "", myerrors.NewCodeError(err)
	}

	return words[1], nil
}

func getNoParam(instructionName string, words []string) *myerrors.CustomError {
	if len(words) != 1 {
		remainingParams := getInstructionParams(words)
		err := myerrors.WrongNumberOfParamsError(instructionName, 0, len(remainingParams), remainingParams)
		return myerrors.NewCodeError(err)
	}
	return nil
}

func getInstructionParams(words []string) []string {
	return words[1:]
}
