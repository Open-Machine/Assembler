package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/core"
	"github.com/open-machine/assembler/helper"
	"gopkg.in/alecthomas/kingpin.v2"
)

func ConfigureRunCommand(app *kingpin.Application) {
	data := &RunCommandConfig{}
	instruction := app.Command("run", "Run machine code").Action(data.run)

	instruction.Arg("file-name", "Provide the name of file with the assembly or machine code").
		Required().
		HintOptions("main.asmm or main.run").
		StringVar(&data.FileName)

	instruction.Arg("number-format", "Provide the format for the numbers printed").
		HintOptions("b (binary), h (hexadeximal) or d (decimal)").
		Default("d").
		StringVar(&data.NumberFormat)
}

type RunCommandConfig struct {
	FileName     string
	NumberFormat string
}

func (data *RunCommandConfig) run(c *kingpin.ParseContext) error {
	extension := helper.FileExtension(data.FileName)
	switch extension {
	case config.AssemblyFileExtension, config.MachineCodeFileExtension:
		var machineCodeFile string
		if extension == config.AssemblyFileExtension {
			file := core.AssembleFile(data.FileName, nil)
			if file == nil {
				break
			} else {
				machineCodeFile = *file
			}

			helper.LogStep(fmt.Sprintf("Running asm file '%s'", data.FileName))
		} else {
			machineCodeFile = data.FileName

			helper.LogStep(fmt.Sprintf("Running assembled file '%s'", data.FileName))
		}

		scriptPath, _ := exec.LookPath(config.RunMachineCodeScriptPath)
		scriptDirectory := scriptPath[:strings.LastIndex(scriptPath, "/")]

		cur_dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		absoluteMachineFile := cur_dir + "/" + machineCodeFile

		cmd := &exec.Cmd{
			Path:   scriptPath,
			Dir:    scriptDirectory,
			Args:   []string{scriptPath, absoluteMachineFile, data.NumberFormat},
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		if cmd == nil {
			helper.LogOtherError("Could not run machine code: running script not found! This is probably a bug, please create an issue report in the repository.")
			return nil
		} else {
			err := cmd.Run()
			if err != nil {
				helper.LogOtherError(fmt.Sprintf("Error: %s.", err))
				helper.LogOtherError("Make sure you have all repository installed in the right configuration!")
			}
		}
	default:
		return fmt.Errorf("unknown extension: please, add a file name with extesion %s or %s", config.AssemblyFileExtension, config.MachineCodeFileExtension)
	}

	return nil
}
