package cli

import (
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
	// TODO: remove
	fmt.Printf("FileName: %v\n", data.FileName)
	fmt.Printf("To Clipboard: %v\n", data.ToClipboard)
	fmt.Printf("Executable filename: %v\n", data.ExecutableFileName)

	// TODO:
	// core.AssembleFile(file)

	return nil
}
