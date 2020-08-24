package cli

import (
	"assembler/helper"
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

// TODO: test

func ConfigureAssembleCommand(app *kingpin.Application) {
	data := &AssembleCommand{}
	command := app.Command("assemble", "Assemble assembly code").Action(data.run)

	command.Arg("file-name", "Provide the name of file with the assembly code").
		Required().
		HintOptions("main.asm").
		StringVar(&data.FileName)

	command.Flag("copy-to-clipboard", "The assembled code will be copied to clipboard instead of written to a file").
		Short('c').
		BoolVar(&data.ToClipboard)

	command.Flag("executable-file-name", "Name of the executable file that will be generated").
		Short('x').
		Default("").
		StringVar(&data.ExecutableFileName)
}

type AssembleCommand struct {
	FileName           string
	ToClipboard        bool
	ExecutableFileName string
}

func (data *AssembleCommand) run(c *kingpin.ParseContext) error {
	helper.PrintOut(fmt.Sprintf("FileName: %v\n", data.FileName))
	helper.PrintOut(fmt.Sprintf("To Clipboard: %v\n", data.ToClipboard))
	helper.PrintOut(fmt.Sprintf("Executable filename: %v\n", data.ExecutableFileName))

	// TODO:
	// core.AssembleFile(file)

	return nil
}
