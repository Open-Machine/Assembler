package data

import (
	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/utils"
)

const (
	VariableParam = iota
	JumpLabelParam
	NoParam
)

type Instruction struct {
	instructionCode  int
	parameter        string
	parameterAddress int
	parameterType    int
}

func NewJumpInstruction(code int, param string) (*Instruction, *myerrors.CustomError) {
	err := utils.CheckJumpLabelName(param)
	if err != nil {
		return nil, myerrors.NewCodeError(err)
	}
	return &Instruction{instructionCode: code, parameter: param, parameterType: JumpLabelParam}, nil
}

func NewVariableInstruction(code int, param string) (*Instruction, *myerrors.CustomError) {
	err := utils.CheckVariableName(param)
	if err != nil {
		return nil, myerrors.NewCodeError(err)
	}
	return &Instruction{instructionCode: code, parameter: param, parameterType: VariableParam}, nil
}

func NewInstructionWithoutParam(code int) Instruction {
	return Instruction{instructionCode: code, parameter: "", parameterType: NoParam, parameterAddress: 0}
}

func NewInstructionTest(code int, name string, n int, t int) *Instruction {
	if !config.Testing {
		return nil
	}
	return &Instruction{code, name, n, t}
}

func (c Instruction) toExecuter() (string, *myerrors.CustomError) {
	str1, err1 := utils.IntToStrHex(c.instructionCode, config.AmntHexDigitsInstruction)
	if err1 != nil {
		return "", myerrors.NewCodeError(err1)
	}

	str2, err2 := utils.IntToStrHex(c.parameterAddress, config.AmntHexDigitsParam)
	if err2 != nil {
		return "", myerrors.NewCodeError(err2)
	}

	return str1 + str2, nil
}
