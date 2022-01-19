package instruction

import "github.com/open-machine/assembler/data"

type InstructionConfig struct {
	code          int
	parameterType int
}

func GetInstructionsConfig() map[string]InstructionConfig {
	return instructions
}

var instructions = map[string]InstructionConfig{
	"nop": {
		code:          0x0,
		parameterType: data.NoParam,
	},
	"copy": {
		code:          0x1,
		parameterType: data.VariableParam,
	},
	"store": {
		code:          0x2,
		parameterType: data.VariableParam,
	},
	"add": {
		code:          0x3,
		parameterType: data.VariableParam,
	},
	"sub": {
		code:          0x4,
		parameterType: data.VariableParam,
	},
	"input": {
		code:          0x7,
		parameterType: data.VariableParam,
	},
	"output": {
		code:          0x8,
		parameterType: data.VariableParam,
	},
	"kill": {
		code:          0x9,
		parameterType: data.NoParam,
	},
	"jmp": {
		code:          0xA,
		parameterType: data.JumpLabelParam,
	},
	"jg": {
		code:          0xB,
		parameterType: data.JumpLabelParam,
	},
	"je": {
		code:          0xD,
		parameterType: data.JumpLabelParam,
	},
	"jl": {
		code:          0xF,
		parameterType: data.JumpLabelParam,
	},
}
