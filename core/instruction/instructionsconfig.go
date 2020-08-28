package instruction

import (
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/data"
)

type InstructionConfig struct {
	code     int
	getParam func(instructionName string, words []string) (*data.InstructionParameter, *myerrors.CustomError)
}

func GetInstructionsConfig() map[string]InstructionConfig {
	return instructions
}

var instructions = map[string]InstructionConfig{
	"nop": InstructionConfig{
		getParam: getParamNoParam,
		code:     0x0,
	},
	"copy": InstructionConfig{
		getParam: getSecondWordAsInt,
		code:     0x1,
	},
	"store": InstructionConfig{
		getParam: getSecondWordAsInt,
		code:     0x2,
	},
	"add": InstructionConfig{
		getParam: getSecondWordAsInt,
		code:     0x3,
	},
	"sub": InstructionConfig{
		getParam: getSecondWordAsInt,
		code:     0x4,
	},
	"input": InstructionConfig{
		getParam: getSecondWordAsInt,
		code:     0x7,
	},
	"output": InstructionConfig{
		getParam: getSecondWordAsInt,
		code:     0x8,
	},
	"kill": InstructionConfig{
		getParam: getParamNoParam,
		code:     0x9,
	},
	"jmp": InstructionConfig{
		getParam: getSecondWordAsIntOrString,
		code:     0xA,
	},
	"jg": InstructionConfig{
		getParam: getSecondWordAsIntOrString,
		code:     0xB,
	},
	"je": InstructionConfig{
		getParam: getSecondWordAsIntOrString,
		code:     0xD,
	},
	"jl": InstructionConfig{
		getParam: getSecondWordAsIntOrString,
		code:     0xF,
	},
}
