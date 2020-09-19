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
	if utils.IsOverflow(uint(code), config.AmntBitsInstruction) {
		err := myerrors.InstructionCodeOverflow(code, config.AmntBitsInstruction)
		return nil, myerrors.NewAssemblerError(err)
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

	str1, err1 := utils.IntToStrHex(c.instructionCode, config.AmntHexDigitsInstruction)
	if err1 != nil {
		return "", myerrors.NewCodeError(err1)
	}

	str2, err2 := utils.IntToStrHex(c.parameter.Num, config.AmntHexDigitsParam)
	if err2 != nil {
		return "", myerrors.NewCodeError(err2)
	}

	return str1 + str2, nil
}
