package commandsassembler

import (
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/data"
)

type CommandAssembler interface {
	Assemble(mapKey string, line string, program *data.Program) *myerrors.CustomError
}
