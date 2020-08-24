package core

import (
	"assembler/data"
	"assembler/myerrors"
	"assembler/utils"
	"strings"
)

func AssembleCommand(line string) (*data.Command, *myerrors.CustomError) {
	arrayWords := strings.Split(line, " ")
	commandName := arrayWords[0]

	commandConfig, exists := commands[commandName]
	if !exists {
		err := myerrors.CommandDoesNotExistError(commandName)
		return nil, myerrors.NewCodeError(err)
	}

	param, paramErr := commandConfig.getParam(commandName, arrayWords)
	if paramErr != nil {
		return nil, paramErr
	}

	commandPointer, customErr := data.NewCommand(commandConfig.code, *param)
	return commandPointer, customErr
}

func getParamNoParam(commandName string, words []string) (*data.CommandParameter, *myerrors.CustomError) {
	if len(words) != 1 {
		remainingParams := getCommandParams(words)
		err := myerrors.WrongNumberOfParamsError(commandName, 0, len(remainingParams), remainingParams)
		return nil, myerrors.NewCodeError(err)
	}

	param := data.NewIntParam(0)
	return &param, nil
}

func getSecondWordAsInt(commandName string, words []string) (*data.CommandParameter, *myerrors.CustomError) {
	return getSecondWord(commandName, words, false)
}

func getSecondWordAsIntOrString(commandName string, words []string) (*data.CommandParameter, *myerrors.CustomError) {
	return getSecondWord(commandName, words, true)
}

func getSecondWord(commandName string, words []string, acceptStringParam bool) (*data.CommandParameter, *myerrors.CustomError) {
	if len(words) != 2 {
		if len(words) < 2 {
			err := myerrors.WrongNumberOfParamsError(commandName, 1, 0, []string{})
			return nil, myerrors.NewCodeError(err)
		}

		remainingParams := getCommandParams(words)
		err := myerrors.WrongNumberOfParamsError(commandName, 1, len(remainingParams), remainingParams)
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

func getCommandParams(words []string) []string {
	return words[1:]
}
