package engine

import (
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/data"
)

func checkAmntWords(amntParamsExpected int, instructionName string, words []string) *myerrors.CustomError {
	params := getInstructionParams(words)
	amntParamsReceived := len(params)
	if amntParamsReceived < amntParamsExpected {
		err := myerrors.WrongNumberOfParamsError(instructionName, amntParamsExpected, amntParamsReceived, params)
		return myerrors.NewCodeError(err)
	} else if amntParamsReceived > amntParamsExpected {
		err := myerrors.WrongNumberOfParamsError(instructionName, amntParamsExpected, amntParamsReceived, params)
		return myerrors.NewCodeError(err)
	}

	return nil
}

func addInstructionStrParam(code int, strParam string, program *data.Program) *myerrors.CustomError {
	param, err := data.NewStringParam(strParam)
	// TODO: test
	if err != nil {
		return myerrors.NewCodeError(err)
	}

	return addInstruction(code, *param, program)
}

func addInstructionIntParam(code int, intParam int, program *data.Program) *myerrors.CustomError {
	param, err := data.NewIntParam(intParam)
	// TODO: test
	if err != nil {
		return myerrors.NewCodeError(err)
	}

	return addInstruction(code, *param, program)
}

func addInstruction(code int, param data.InstructionParameter, program *data.Program) *myerrors.CustomError {
	instruction, errNew := data.NewInstruction(code, param)
	if errNew != nil {
		return myerrors.NewCodeError(errNew)
	}

	program.AddInstruction(*instruction)
	return nil
}

func getInstructionParams(words []string) []string {
	return words[1:]
}
