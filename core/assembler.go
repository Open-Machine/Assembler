package core

import (
	"fmt"
	"os"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/myerrors"
	"github.com/open-machine/assembler/utils/helper"
)

func AssembleFile(path string, execFileNameParam *string) {
	status, execFileName := assembleFile(path, execFileNameParam)
	if status == config.FailStatus {
		helper.LogOtherError(fmt.Sprintf("========= Failed to assemble %s =========", path))
	} else {
		helper.LogInfo(fmt.Sprintf("========= Executable file is available in '%s' =========", execFileName))
	}
}

func assembleFile(path string, execFileNameParam *string) (int, string) {
	if helper.FileExtension(path) != config.AssemblyFileExtension {
		helper.LogOtherError(fmt.Sprintf("Invalid extension: assembly file should have '%s' extension.", config.AssemblyFileExtension))
		return config.FailStatus, ""
	}

	file, err := os.Open(path)
	if err != nil {
		helper.LogOtherError(fmt.Sprintf("Cannot open file '%s'. Error: %s", path, err.Error()))
		return config.FailStatus, ""
	}
	defer file.Close()

	helper.LogInfo(fmt.Sprintf("========= Starting to assemble '%s' =========", path))

	ptrProgram := programFromFile(file)
	if ptrProgram == nil {
		return config.FailStatus, ""
	}

	errs := ptrProgram.ReplaceLabelsWithNumbers()
	if len(errs) > 0 {
		for _, err := range errs {
			// TODO: create infrastructure go get lineIndex and line here
			helper.LogErrorInLine(*myerrors.NewCodeError(err), 0, "")
		}
		return config.FailStatus, ""
	}

	var execFileName string
	if execFileNameParam == nil {
		execFileName = helper.FileNameWithoutExtension(file.Name())
	} else {
		execFileName = *execFileNameParam
	}

	execFile, err := os.Create(execFileName)
	if err != nil {
		helper.LogOtherError(fmt.Sprintf("Cannot open file '%s'. Error: %s", execFileName, err.Error()))
		return config.FailStatus, ""
	}

	writeStatus := writeExecProgram(*ptrProgram, execFileName, execFile)
	if writeStatus == config.FailStatus {
		return config.FailStatus, ""
	}

	return config.SuccessStatus, execFileName
}
