package cli

import (
	"strings"

	"github.com/open-machine/assembler/core"

	"gopkg.in/alecthomas/kingpin.v2"
)

func ConfigureAssembleCommand(app *kingpin.Application) {
	data := &AssembleInstruction{}
	instruction := app.Command("assemble", "Assemble assembly code").Action(data.run)

	instruction.Arg("file-name", "Provide the name of file with the assembly code").
		Required().
		HintOptions("main.asm").
		StringVar(&data.FileName)

	instruction.Flag("rename-exec", "Provide the name of the executable file that will be created (if empty, the name will be the same as the assembly code file)").
		Short('r').
		Default("").
		StringVar(&data.ExecutableFileName)
}

type AssembleInstruction struct {
	FileName           string
	ExecutableFileName string
}

func (data *AssembleInstruction) run(c *kingpin.ParseContext) error {
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
