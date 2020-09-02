package data

import (
	"github.com/open-machine/assembler/config/myerrors"
)

type Program struct {
	instructions   []Instruction
	jumpLabelsDict map[string]int
}

func NewProgram(lines int) Program {
	return Program{instructions: make([]Instruction, 0, lines), jumpLabelsDict: map[string]int{}}
}

func ProgramFromInstructionsAndLabels(instructions []Instruction, jumpLabelsDict map[string]int) Program {
	return Program{instructions: instructions, jumpLabelsDict: jumpLabelsDict}
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

func (p *Program) ReplaceLabelsWithNumbers() []error {
	errs := make([]error, 0)

	for i, instruction := range p.instructions {
		if instruction.parameter.IsStr {
			label := instruction.parameter.Str
			instructionIndex, exists := p.jumpLabelsDict[label]

			if !exists {
				errs = append(errs, myerrors.JumpLabelDoesNotExistError(label))
			} else {
				instruction.parameter = NewIntParam(instructionIndex)
				p.instructions[i] = instruction
			}
		}
	}

	if len(errs) == 0 {
		p.jumpLabelsDict = map[string]int{}
	}
	return errs
}

const executerFileHeader = "v2.0 raw"

func (p *Program) ToExecuter() (string, []myerrors.CustomError) {
	machineCodeStr := executerFileHeader + "\n" + "0000"
	// TODO: Circuit requires "0000" as first (probably because of the PC counter and inverted clock)

	var errors []myerrors.CustomError

	for _, instruction := range p.instructions {
		executerCode, err := instruction.toExecuter()
		if err != nil {
			errors = append(errors, *err)
		}

		machineCodeStr += executerCode
	}

	if len(errors) == 0 {
		return machineCodeStr, errors
	}
	return "", errors
}
