package cli

import (
	"assembler/core"
	"assembler/helper"
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

// TODO: tests

func ConfigureSyntaxCommand(app *kingpin.Application) {
	data := &SyntaxCommand{}
	command := app.Command("syntax", "Help with the syntax of this assembly language").Action(data.run)

	command.Flag("example", "Assembly code example with explanation").
		Short('e').
		BoolVar(&data.Example)

	command.Flag("ls", "List all available commands").
		Short('l').
		BoolVar(&data.Ls)

	command.Flag("command", "Explanation of an specific command").
		Short('c').
		HintAction(listCommands).
		EnumVar(&data.Command, listCommands()...)
}

type SyntaxCommand struct {
	Example bool
	Ls      bool
	Command string
}

func (data *SyntaxCommand) run(c *kingpin.ParseContext) error {
	var err error = nil

	if data.Example {
		// TODO
	}

	if data.Ls {
		if data.Example {
			helper.PrintlnExplanation("")
		}

		helper.PrintlnExplanation("Assembly commands list and explanations:")
		helper.PrintlnExplanation("")
		sortedCmdNames, cmdExplanations := core.GetCommandsExplanationSorted()

		for _, name := range sortedCmdNames {
			explanation := cmdExplanations[name]

			helper.PrintlnExplanation(fmt.Sprintf("   - %s:\t%s", name, explanation.Command))
			helper.PrintlnExplanation("")
		}
	}

	if data.Command != "" {
		if data.Ls {
			helper.PrintlnExplanation("")
		}

		explanation, _ := core.GetCommandExplanation(data.Command)
		// kingpin won't let the request go throught if the command name does not exist (because of the enum)

		helper.PrintlnExplanation(fmt.Sprintf("'%s' command explanation:", data.Command))
		helper.PrintlnExplanation(fmt.Sprintf("\t\tCommand: %s", explanation.Command))
		helper.PrintlnExplanation(fmt.Sprintf("\t\tParameter: %s", explanation.Param))
	}

	return err
}

func listCommands() []string {
	commands := make([]string, 0)
	for name := range core.GetCommandsExplanation() {
		commands = append(commands, name)
	}

	return commands
}

func getSyntaxExample() []string {
	return []string{}
}
