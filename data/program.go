package data

import (
	"assembler/myerrors"
)

type Program struct {
	commands       []Command
	jumpLabelsDict map[string]int
}

func NewProgram(lines int) Program {
	return Program{commands: make([]Command, 0, lines), jumpLabelsDict: map[string]int{}}
}

func ProgramFromCommandsAndLabels(commands []Command, jumpLabelsDict map[string]int) Program {
	return Program{commands: commands, jumpLabelsDict: jumpLabelsDict}
}

func (p *Program) AddCommand(command Command) {
	p.commands = append(p.commands, command)
}

func (p *Program) LenCommands() int {
	return len(p.commands)
}

func (p *Program) AddJumpLabel(label string, commandIndex int) error {
	_, exists := p.jumpLabelsDict[label]
	if exists {
		return myerrors.JumpLabelAlreadyExistsError(label)
	}

	p.jumpLabelsDict[label] = commandIndex
	return nil
}

func (p *Program) ReplaceLabelsWithNumbers() []error {
	errs := make([]error, 0)

	for i, command := range p.commands {
		if command.parameter.IsStr {
			label := command.parameter.Str
			commandIndex, exists := p.jumpLabelsDict[label]

			if !exists {
				errs = append(errs, myerrors.JumpLabelDoesNotExistError(label))
			} else {
				command.parameter = NewIntParam(commandIndex)
				p.commands[i] = command
			}
		}
	}

	if len(errs) == 0 {
		p.jumpLabelsDict = map[string]int{}
	}
	return errs
}

func (p *Program) ToExecuter() (string, []myerrors.CustomError) {
	str := ""
	var errors []myerrors.CustomError

	for _, command := range p.commands {
		executerCode, err := command.toExecuter()
		if err != nil {
			errors = append(errors, *err)
		}

		str += executerCode
	}

	return str, errors
}
