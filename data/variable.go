package data

import (
	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/utils"
)

type Variable struct {
	name         string
	index        int
	initialValue uint
}

func NewVariable(name string, index int, initialValue uint) (*Variable, *myerrors.CustomError) {
	err := utils.CheckVariableName(name)
	if err != nil {
		return nil, myerrors.NewCodeError(err)
	}

	if utils.IsOverflow(initialValue, config.AmntBitsVariable) {
		return nil, myerrors.NewCodeError(myerrors.ValueOverflow(initialValue, config.AmntBitsVariable))
	}
	return &Variable{name: name, index: index, initialValue: initialValue}, nil
}
