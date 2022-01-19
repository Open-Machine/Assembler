package cli

import (
	"fmt"

	"github.com/open-machine/assembler/config/instructionsexplanation"
	"github.com/open-machine/assembler/helper"

	"gopkg.in/alecthomas/kingpin.v2"
)

func ConfigureSyntaxCommand(app *kingpin.Application) {
	data := &SyntaxInstruction{}
	instruction := app.Command("syntax", "Help with the syntax of this assembly language").Action(data.run)

	instruction.Flag("example", "Assembly code example with explanation").
		Short('e').
		BoolVar(&data.Example)

	instruction.Flag("ls", "List all available instructions").
		Short('l').
		BoolVar(&data.Ls)

	instruction.Flag("instruction", "Explanation of an specific instruction").
		Short('c').
		HintAction(listInstructions).
		EnumVar(&data.Instruction, listInstructions()...)
}

type SyntaxInstruction struct {
	Example     bool
	Ls          bool
	Instruction string
}

func (data *SyntaxInstruction) run(c *kingpin.ParseContext) error {
	var err error = nil

	if data.Example {
		helper.PrintlnExplanation("# Assembly Example")
		lines := GetAssemblyExample()
		for _, line := range lines {
			helper.PrintlnExplanation(line)
		}
	}

	if data.Ls {
		if data.Example {
			helper.PrintlnExplanation("")
		}

		helper.PrintlnExplanation("Assembly instructions list and explanations:")
		helper.PrintlnExplanation("")
		sortedCmdNames, cmdExplanations := instructionsexplanation.GetInstructionsExplanationSorted()

		for _, name := range sortedCmdNames {
			explanation := cmdExplanations[name]

			helper.PrintlnExplanation(fmt.Sprintf("   - %s:\t%s", name, explanation.Instruction))
			helper.PrintlnExplanation("")
		}
	}

	if data.Instruction != "" {
		if data.Ls {
			helper.PrintlnExplanation("")
		}

		explanation, _ := instructionsexplanation.GetInstructionExplanation(data.Instruction)
		// kingpin won't let the request go throught if the instruction name does not exist (because of the enum)

		helper.PrintlnExplanation(fmt.Sprintf("'%s' instruction explanation:", data.Instruction))
		helper.PrintlnExplanation(fmt.Sprintf("\t\tInstruction: %s", explanation.Instruction))
		helper.PrintlnExplanation(fmt.Sprintf("\t\tParameter: %s", explanation.Param))
	}

	return err
}

func GetAssemblyExample() []string {
	return []string{
		"@VAR",
		"a = 10",
		"b = 20",
		"",
		"@CODE",
		"copy a",
		"add b",
		"store b",
		"kill",
	}
}

func listInstructions() []string {
	list := make([]string, 0)
	for name := range instructionsexplanation.GetInstructionsExplanation() {
		list = append(list, name)
	}

	return list
}
