package commandsassembler

import (
	"github.com/open-machine/assembler/config/myerrors"
)

type AssemblerConfig struct {
	mapAssemblers   map[string]CommandAssembler
	regexAssemblers []RegexCommandAssembler
	last            *CommandAssembler
}

func NewAssemblerConfig() AssemblerConfig {
	return AssemblerConfig{mapAssemblers: map[string]CommandAssembler{}, regexAssemblers: []RegexCommandAssembler{}, last: nil}
}

// TODO: check every command if it is only accepted in only one regex (in the beginning test if the number of tests is equal to the setup)

func (a *AssemblerConfig) AppendRegexAssembler(assembler RegexCommandAssembler) *myerrors.CustomError {
	a.regexAssemblers = append(a.regexAssemblers, assembler)
	return nil
}

func (a *AssemblerConfig) AppendMapAssembler(key string, assembler CommandAssembler) *myerrors.CustomError {
	_, exists := a.mapAssemblers[key]
	if exists {
		return myerrors.NewAssemblerError(myerrors.CommandAlreadyExistError(key))
	}

	a.mapAssemblers[key] = assembler
	return nil
}
