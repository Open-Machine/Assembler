package data

import (
	"fmt"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/utils"
)

type Program struct {
	instructions   []Instruction
	jumpLabelsDict map[string]int
	variablesDict  map[string]Variable
}

func NewProgram(lines int) Program {
	return Program{instructions: make([]Instruction, 0, lines), jumpLabelsDict: map[string]int{}, variablesDict: map[string]Variable{}}
}

func NewCompleteProgram(instructions []Instruction, jumpLabelsDict map[string]int, variablesDict map[string]Variable) Program {
	return Program{instructions: instructions, jumpLabelsDict: jumpLabelsDict, variablesDict: variablesDict}
}

func (p *Program) AddInstruction(instruction Instruction) {
	p.instructions = append(p.instructions, instruction)
}

func (p *Program) LenInstructions() int {
	return len(p.instructions)
}

func (p *Program) AddJumpLabel(label string, instructionIndex int) error {
	_, exists := p.jumpLabelsDict[label]
	if exists {
		return myerrors.JumpLabelAlreadyExistsError(label)
	}

	p.jumpLabelsDict[label] = instructionIndex
	return nil
}

func (p *Program) AddVariable(variable Variable) error {
	_, exists := p.variablesDict[variable.name]
	if exists {
		return myerrors.VariableAlreadyExistsError(variable.name)
	}

	p.variablesDict[variable.name] = variable
	return nil
}

func (p *Program) ReplaceLabelsWithAddresses() []error {
	errs := make([]error, 0)
	for i, instruction := range p.instructions {
		switch instruction.parameterType {
		case VariableParam:
			variable, exists := p.variablesDict[instruction.parameter]
			if !exists {
				errs = append(errs, myerrors.VariableDoesNotExistError(instruction.parameter))
			} else {
				p.instructions[i].parameterAddress = config.MemorySlots - variable.index - 1
			}

		case JumpLabelParam:
			instructionIndex, exists := p.jumpLabelsDict[instruction.parameter]
			if !exists {
				errs = append(errs, myerrors.JumpLabelDoesNotExistError(instruction.parameter))
			} else {
				p.instructions[i].parameterAddress = instructionIndex
			}

		case NoParam:
		}
	}

	return errs
}

const executerFileHeader = "v2.0 raw"

func (p *Program) ToExecuter() (string, []myerrors.CustomError) {
	machineCodeStr := executerFileHeader + "\n" + "0000" + " "
	// TODO: Circuit requires "0000" as first (probably because of the PC counter and inverted clock)

	var errors []myerrors.CustomError

	for _, instruction := range p.instructions {
		executerCode, err := instruction.toExecuter()
		if err != nil {
			errors = append(errors, *err)
		}
		machineCodeStr += executerCode + " "
	}

	amountEmpty := config.MemorySlots - len(p.instructions) - len(p.variablesDict) - 1
	machineCodeStr += fmt.Sprintf("%d*0 ", amountEmpty)

	variablesInOrder := make([]Variable, len(p.variablesDict))
	for varname := range p.variablesDict {
		variable := p.variablesDict[varname]
		index := len(p.variablesDict) - variable.index - 1 // from back to front
		variablesInOrder[index] = variable
	}
	for _, variable := range variablesInOrder {
		initialValueStr, err := utils.IntToStrHex(
			int(variable.initialValue),
			config.AmntHexDigitsVariable,
		)
		if err != nil {
			errors = append(errors, *myerrors.NewCodeError(err))
		}
		machineCodeStr += initialValueStr + " "
	}

	if len(errors) == 0 {
		return machineCodeStr, errors
	}
	return "", errors
}
