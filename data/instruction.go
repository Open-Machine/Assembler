package data

import (
	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/utils"
)

type Instruction struct {
	instructionCode int
	parameter       InstructionParameter
}

func NewInstruction(code int, param InstructionParameter) (*Instruction, *myerrors.CustomError) {
	if utils.IsOverflow(uint(code), config.AmntBitsCode) {
		err := myerrors.InstructionCodeOverflow(code, config.AmntBitsCode)
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

	return &Instruction{code, param}, nil
}

func NewInstructionTest(code int, param InstructionParameter) *Instruction {
	if !config.Testing {
		return nil
	}
	return &Instruction{code, param}
}

func (c Instruction) toExecuter() (string, *myerrors.CustomError) {
	if c.parameter.IsStr {
		return "", myerrors.NewAssemblerError(myerrors.InvalidStateTransformationToExecuterError())
	}

	str1, err1 := utils.IntToStrHex(c.instructionCode, 2)
	if err1 != nil {
		return "", myerrors.NewCodeError(err1)
	}

	str2, err2 := utils.IntToStrHex(c.parameter.Num, 2)
	if err2 != nil {
		return "", myerrors.NewCodeError(err2)
	}

	return str1 + str2, nil
}
