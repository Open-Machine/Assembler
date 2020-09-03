package commandsassembler

import "github.com/open-machine/assembler/data"

type funcAssemble func(line string, program *data.Program) (*data.Program, error)

type CommandAssembler struct {
	Assemble funcAssemble
}

func NewCommandAssembler(assemble funcAssemble) CommandAssembler {
	return CommandAssembler{Assemble: assemble}
}
