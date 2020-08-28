package data

import (
	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/myerrors"
	"github.com/open-machine/assembler/utils"
)

type Command struct {
	commandCode int
	parameter   CommandParameter
}

func NewCommand(code int, param CommandParameter) (*Command, *myerrors.CustomError) {
	if utils.IsOverflow(uint(code), config.AmntBitsCode) {
		err := myerrors.CommandCodeOverflow(code, config.AmntBitsCode)
		return nil, myerrors.NewAssemblerError(err)
	}

	if param.IsStr {
		if !utils.IsValidVarName(param.Str) {
			err := myerrors.InvalidLabelParam(param.Str)
			return nil, myerrors.NewAssemblerError(err)
		}
	} else {
		if !param.IsStr && utils.IsOverflow(uint(param.Num), config.AmntBitsParam) {
			err := myerrors.ParamOverflow(param.Num, config.AmntBitsParam)
			return nil, myerrors.NewCodeError(err)
		}
	}

	return &Command{code, param}, nil
}

func NewCommandTest(code int, param CommandParameter) *Command {
	if !config.Testing {
		return nil
	}
	return &Command{code, param}
}

func (c Command) toExecuter() (string, *myerrors.CustomError) {
	if c.parameter.IsStr {
		return "", myerrors.NewAssemblerError(myerrors.InvalidStateTransformationToExecuterError())
	}

	str1, err1 := utils.IntToStrHex(c.commandCode, 2)
	if err1 != nil {
		return "", myerrors.NewCodeError(err1)
	}

	str2, err2 := utils.IntToStrHex(c.parameter.Num, 2)
	if err2 != nil {
		return "", myerrors.NewCodeError(err2)
	}

	return str1 + str2, nil
}
