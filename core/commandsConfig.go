package core

import (
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/myerrors"
	"sort"
)

type commandConfig struct {
	code               int
	getParam           func(commandName string, words []string) (*data.CommandParameter, *myerrors.CustomError)
	commandExplanation string
	paramExplanation   string
}

var commands = map[string]commandConfig{
	"nop": commandConfig{
		// Config
		getParam: getParamNoParam,
		code:     0x0,
		// Syntax explanation
		commandExplanation: "No operation",
		paramExplanation:   "No param needed",
	},
	"copy": commandConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x1,
		// Syntax explanation
		commandExplanation: "Copies the value from memory to the ACC register",
		paramExplanation:   "The parameter refers to the memory address",
	},
	"store": commandConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x2,
		// Syntax explanation
		commandExplanation: "Stores the value of the ACC register in memory",
		paramExplanation:   "The parameter refers to the memory address",
	},
	"add": commandConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x3,
		// Syntax explanation
		commandExplanation: "Adds a memory value to the ACC register and stores the result in ACC",
		paramExplanation:   "The parameter refers to the memory address",
	},
	"sub": commandConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x4,
		// Syntax explanation
		commandExplanation: "Subtracts a memory value from the value of the ACC register and stores the result in ACC",
		paramExplanation:   "The parameter refers to the memory address",
	},
	"input": commandConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x7,
		// Syntax explanation
		commandExplanation: "Inputs the input value into memory",
		paramExplanation:   "The parameter refers to the memory address",
	},
	"output": commandConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x8,
		// Syntax explanation
		commandExplanation: "Outputs a memory value",
		paramExplanation:   "The parameter refers to the memory address",
	},
	"kill": commandConfig{
		// Config
		getParam: getParamNoParam,
		code:     0x9,
		// Syntax explanation
		commandExplanation: "Kills the program (you will need this command to tell the computer that the program ended)",
		paramExplanation:   "No param needed",
	},
	"jmp": commandConfig{
		// Config
		getParam: getSecondWordAsIntOrString,
		code:     0xA,
		// Syntax explanation
		commandExplanation: "Jumps to the a command",
		paramExplanation:   "The parameter can be either a label or a number that refers to a command (Warning: the index of a command can be different from the index of the line where the command is)",
	},
	"jg": commandConfig{
		// Config
		getParam: getSecondWordAsIntOrString,
		code:     0xB,
		// Syntax explanation
		commandExplanation: "Jumps to the a command if ACC register is greater than zero",
		paramExplanation:   "The parameter can be either a label or a number that refers to a command (Warning: the index of a command can be different from the index of the line where the command is)",
	},
	"je": commandConfig{
		// Config
		getParam: getSecondWordAsIntOrString,
		code:     0xD,
		// Syntax explanation
		commandExplanation: "Jumps to the a command if ACC register is equal to zero",
		paramExplanation:   "The parameter can be either a label or a number that refers to a command (Warning: the index of a command can be different from the index of the line where the command is)",
	},
	"jl": commandConfig{
		// Config
		getParam: getSecondWordAsIntOrString,
		code:     0xF,
		// Syntax explanation
		commandExplanation: "Jumps to the a command if ACC register is less than zero",
		paramExplanation:   "The parameter can be either a label or a number that refers to a command (Warning: the index of a command can be different from the index of the line where the command is)",
	},
}

type CommandExplanation struct {
	Command string
	Param   string
}

func newCommandExplanation(cmdConfig commandConfig) CommandExplanation {
	return CommandExplanation{Command: cmdConfig.commandExplanation, Param: cmdConfig.paramExplanation}
}

func GetCommandsExplanation() map[string]CommandExplanation {
	cmdExplanations := map[string]CommandExplanation{}
	for name, config := range commands {
		cmdExplanations[name] = newCommandExplanation(config)
	}
	return cmdExplanations
}

func GetCommandsExplanationSorted() ([]string, map[string]CommandExplanation) {
	explanations := GetCommandsExplanation()

	keys := make([]string, len(explanations))
	i := 0
	for k := range explanations {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	return keys, explanations
}

func GetCommandExplanation(name string) (*CommandExplanation, error) {
	config, exists := commands[name]
	if !exists {
		return nil, myerrors.CommandDoesNotExistError(name)
	}
	explanation := newCommandExplanation(config)
	return &explanation, nil
}
