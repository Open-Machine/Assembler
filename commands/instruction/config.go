package instruction

import (
	"github.com/open-machine/assembler/commands/instruction/engine"
	"github.com/open-machine/assembler/core/commandsassembler"
)

func GetInstructionsConfig() map[string]commandsassembler.CommandAssembler {
	return instructions
}

// TODO: test if any instruction has the same code
var instructions = map[string]commandsassembler.CommandAssembler{
	"nop":    engine.NewNoParamInstruct(0x0),
	"copy":   engine.NewOneParamInstruct(0x1, !engine.ACCEPT_STRING_PARAM),
	"store":  engine.NewOneParamInstruct(0x2, !engine.ACCEPT_STRING_PARAM),
	"add":    engine.NewOneParamInstruct(0x3, !engine.ACCEPT_STRING_PARAM),
	"sub":    engine.NewOneParamInstruct(0x4, !engine.ACCEPT_STRING_PARAM),
	"input":  engine.NewOneParamInstruct(0x7, !engine.ACCEPT_STRING_PARAM),
	"output": engine.NewOneParamInstruct(0x8, !engine.ACCEPT_STRING_PARAM),
	"kill":   engine.NewNoParamInstruct(0x9),
	"jmp":    engine.NewOneParamInstruct(0xA, engine.ACCEPT_STRING_PARAM),
	"jg":     engine.NewOneParamInstruct(0xB, engine.ACCEPT_STRING_PARAM),
	"je":     engine.NewOneParamInstruct(0xD, engine.ACCEPT_STRING_PARAM),
	"jl":     engine.NewOneParamInstruct(0xF, engine.ACCEPT_STRING_PARAM),
}
