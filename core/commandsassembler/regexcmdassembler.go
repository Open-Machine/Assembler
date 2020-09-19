package commandsassembler

import (
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/data"
)

type RegexCommandAssembler struct {
	regex string
	fun   func(mapKey string, line string, program *data.Program) *myerrors.CustomError
}

func NewRegexCommandAssembler(regex string, f func(mapKey string, line string, program *data.Program) *myerrors.CustomError) RegexCommandAssembler {
	return RegexCommandAssembler{regex: regex, fun: f}
}

func RegexAssemblerFrom(regex string, assembler CommandAssembler) RegexCommandAssembler {
	return RegexCommandAssembler{regex: regex, fun: assembler.Assemble}
}

func (regexassemble *RegexCommandAssembler) Assemble(k string, l string, p *data.Program) *myerrors.CustomError {
	return regexassemble.fun(k, l, p)
}
