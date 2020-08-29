package core

import (
	"fmt"
	"io"
	"os"

	"github.com/open-machine/assembler/utils"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/helper"
)

func AssembleFile(path string, execFileNameParam *string) {
	helper.LogInfo(fmt.Sprintf("========= Starting to assemble '%s' =========", path))

	status, execFileName := AssembleFileAux(path, execFileNameParam, ioReaderFromPath, ioWriterFromPath)

	if status == config.FailStatus {
		helper.LogOtherError(fmt.Sprintf("========= Failed to assemble %s =========", path))
	} else {
		helper.LogInfo(fmt.Sprintf("========= Executable file is available in '%s' =========", execFileName))
	}
}

func ioReaderFromPath(path string) (utils.MyFileInterface, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	myNewFile := utils.NewMyFile(*file)
	return &myNewFile, nil
}

func ioWriterFromPath(path string) (io.Writer, error) {
	return os.Create(path)
}

type ioReaderFromPathFun func(string) (utils.MyFileInterface, error)
type ioWriterFromPathFun func(string) (io.Writer, error)

func AssembleFileAux(path string, execFileNameParam *string, ioReaderFromPath ioReaderFromPathFun, ioWriterFromPath ioWriterFromPathFun) (int, string) {
	if helper.FileExtension(path) != config.AssemblyFileExtension {
		helper.LogOtherError(fmt.Sprintf("Invalid extension: assembly file should have '%s' extension.", config.AssemblyFileExtension))
		return config.FailStatus, ""
	}

	file, err := ioReaderFromPath(path)
	if err != nil {
		helper.LogOtherError(fmt.Sprintf("Cannot open file '%s'. Error: %s", path, err.Error()))
		return config.FailStatus, ""
	}
	defer file.Close()

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

	// TODO: verify execFileName as possible name or not ([a-zA-Z_-]+)

	var execFileName string
	if execFileNameParam == nil {
		execFileName = helper.FileNameWithoutExtension(file.Name())
	} else {
		execFileName = *execFileNameParam
	}

	execFile, err := ioWriterFromPath(execFileName)
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
