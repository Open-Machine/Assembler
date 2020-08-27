package cli

import (
	"assembler/core"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

func ConfigureAssembleCommand(app *kingpin.Application) {
	data := &AssembleCommand{}
	command := app.Command("assemble", "Assemble assembly code").Action(data.run)

	command.Arg("file-name", "Provide the name of file with the assembly code").
		Required().
		HintOptions("main.asm").
		StringVar(&data.FileName)

	command.Flag("rename-exec", "Provide the name of the executable file that will be created (if empty, the name will be the same as the assembly code file)").
		Short('r').
		Default("").
		StringVar(&data.ExecutableFileName)
}

type AssembleCommand struct {
	FileName           string
	ExecutableFileName string
}

func (data *AssembleCommand) run(c *kingpin.ParseContext) error {
	data.ExecutableFileName = strings.TrimSpace(data.ExecutableFileName)

	var execFileNameParam *string
	if data.ExecutableFileName == "" {
		execFileNameParam = nil
	} else {
		execFileNameParam = &data.ExecutableFileName
	}

	core.AssembleFile(data.FileName, execFileNameParam)
	return nil
}
