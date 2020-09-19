package data

import (
	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/utils"
)

type InstructionParameter struct {
	Num   int
	Str   string
	IsStr bool
}

func NewStringParam(str string) (*InstructionParameter, *myerrors.CustomError) {
	err := utils.CheckParamName(str)
	if err != nil {
		return nil, myerrors.NewAssemblerError(err)
	}

	return &InstructionParameter{Num: 0, Str: str, IsStr: true}, nil
}

func NewIntParam(num int) (*InstructionParameter, *myerrors.CustomError) {
	if utils.IsOverflow(uint(num), config.AmntBitsParam) {
		err := myerrors.ParamOverflow(num, config.AmntBitsParam)
		return nil, myerrors.NewCodeError(err)
	}

	return &InstructionParameter{Num: num, Str: "", IsStr: false}, nil
}

func NewParamTest(num int, str string, isStr bool) *InstructionParameter {
	if config.Testing {
		return &InstructionParameter{Num: num, Str: str, IsStr: isStr}
	} else {
		return nil
	}
}
