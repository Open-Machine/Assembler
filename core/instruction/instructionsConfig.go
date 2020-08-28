package instruction

import (
	"sort"

	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/myerrors"
)

type instructionConfig struct {
	code                   int
	getParam               func(instructionName string, words []string) (*data.InstructionParameter, *myerrors.CustomError)
	instructionExplanation string
	paramExplanation       string
}

var instructions = map[string]instructionConfig{
	"nop": instructionConfig{
		// Config
		getParam: getParamNoParam,
		code:     0x0,
		// Syntax explanation
		instructionExplanation: "No operation",
		paramExplanation:       "No param needed",
	},
	"copy": instructionConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x1,
		// Syntax explanation
		instructionExplanation: "Copies the value from memory to the ACC register",
		paramExplanation:       "The parameter refers to the memory address",
	},
	"store": instructionConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x2,
		// Syntax explanation
		instructionExplanation: "Stores the value of the ACC register in memory",
		paramExplanation:       "The parameter refers to the memory address",
	},
	"add": instructionConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x3,
		// Syntax explanation
		instructionExplanation: "Adds a memory value to the ACC register and stores the result in ACC",
		paramExplanation:       "The parameter refers to the memory address",
	},
	"sub": instructionConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x4,
		// Syntax explanation
		instructionExplanation: "Subtracts a memory value from the value of the ACC register and stores the result in ACC",
		paramExplanation:       "The parameter refers to the memory address",
	},
	"input": instructionConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x7,
		// Syntax explanation
		instructionExplanation: "Inputs the input value into memory",
		paramExplanation:       "The parameter refers to the memory address",
	},
	"output": instructionConfig{
		// Config
		getParam: getSecondWordAsInt,
		code:     0x8,
		// Syntax explanation
		instructionExplanation: "Outputs a memory value",
		paramExplanation:       "The parameter refers to the memory address",
	},
	"kill": instructionConfig{
		// Config
		getParam: getParamNoParam,
		code:     0x9,
		// Syntax explanation
		instructionExplanation: "Kills the program (you will need this instruction to tell the computer that the program ended)",
		paramExplanation:       "No param needed",
	},
	"jmp": instructionConfig{
		// Config
		getParam: getSecondWordAsIntOrString,
		code:     0xA,
		// Syntax explanation
		instructionExplanation: "Jumps to the a instruction",
		paramExplanation:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
	"jg": instructionConfig{
		// Config
		getParam: getSecondWordAsIntOrString,
		code:     0xB,
		// Syntax explanation
		instructionExplanation: "Jumps to the a instruction if ACC register is greater than zero",
		paramExplanation:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
	"je": instructionConfig{
		// Config
		getParam: getSecondWordAsIntOrString,
		code:     0xD,
		// Syntax explanation
		instructionExplanation: "Jumps to the a instruction if ACC register is equal to zero",
		paramExplanation:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
	"jl": instructionConfig{
		// Config
		getParam: getSecondWordAsIntOrString,
		code:     0xF,
		// Syntax explanation
		instructionExplanation: "Jumps to the a instruction if ACC register is less than zero",
		paramExplanation:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
}

type InstructionExplanation struct {
	Instruction string
	Param       string
}

func newInstructionExplanation(cmdConfig instructionConfig) InstructionExplanation {
	return InstructionExplanation{Instruction: cmdConfig.instructionExplanation, Param: cmdConfig.paramExplanation}
}

func GetInstructionsExplanation() map[string]InstructionExplanation {
	cmdExplanations := map[string]InstructionExplanation{}
	for name, config := range instructions {
		cmdExplanations[name] = newInstructionExplanation(config)
	}
	return cmdExplanations
}

func GetInstructionsExplanationSorted() ([]string, map[string]InstructionExplanation) {
	explanations := GetInstructionsExplanation()

	keys := make([]string, len(explanations))
	i := 0
	for k := range explanations {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	return keys, explanations
}

func GetInstructionExplanation(name string) (*InstructionExplanation, error) {
	config, exists := instructions[name]
	if !exists {
		return nil, myerrors.InstructionDoesNotExistError(name)
	}
	explanation := newInstructionExplanation(config)
	return &explanation, nil
}
