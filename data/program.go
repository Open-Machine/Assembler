package data

import (
	"assembler/myerrors"
)

type Program struct {
	commands       []Command
	gotoLabelsDict map[string]int
}

func NewProgram(lines int) Program {
	return Program{commands: make([]Command, 0, lines), gotoLabelsDict: map[string]int{}}
}

func ProgramFromCommandsAndLabels(commands []Command, gotoLabelsDict map[string]int) Program {
	return Program{commands: commands, gotoLabelsDict: gotoLabelsDict}
}

func (p *Program) AddCommand(command Command) {
	p.commands = append(p.commands, command)
}

func (p *Program) LenCommands() int {
	return len(p.commands)
}

func (p *Program) AddGotoLabel(label string, commandIndex int) error {
	_, exists := p.gotoLabelsDict[label]
	if exists {
		return myerrors.GotoLabelAlreadyExistsError(label)
	}

	p.gotoLabelsDict[label] = commandIndex
	return nil
}

func (p *Program) ReplaceLabelsWithNumbers() []error {
	errs := make([]error, 0)

	for i, command := range p.commands {
		if command.parameter.IsStr {
			label := command.parameter.Str
			commandIndex, exists := p.gotoLabelsDict[label]

			if !exists {
				errs = append(errs, myerrors.GotoLabelDoesNotExistError(label))
			} else {
				command.parameter = NewIntParam(commandIndex)
				p.commands[i] = command
			}
		}
	}

	if len(errs) == 0 {
		p.gotoLabelsDict = map[string]int{}
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
