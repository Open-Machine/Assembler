package core

import (
	"assembler/config"
	"assembler/data"
	"assembler/utils/helper"
	"assembler/myerrors"
	"assembler/utils"
	"bufio"
	"fmt"
	"io"
	"os"
)

const Success = 0
const Fail = 1

func AssembleFile(path string, execFileNameParam *string) {
	status, execFileName := assembleFile(path, execFileNameParam)
	if status == Fail {
		helper.LogOtherError(fmt.Sprintf("========= Failed to assemble %s =========", path))
	} else {
		helper.LogInfo(fmt.Sprintf("========= Executable file is available in '%s' =========", execFileName))
	}
}

func assembleFile(path string, execFileNameParam *string) (int, string) {
	if helper.FileExtension(path) != config.AssemblyFileExtension {
		helper.LogOtherError(fmt.Sprintf("Invalid extension: assembly file should have '%s' extension.", config.AssemblyFileExtension))
		return Fail, ""
	}

	file, err := os.Open(path)
	if err != nil {
		helper.LogOtherError(fmt.Sprintf("Cannot open file '%s'. Error: %s", path, err.Error()))
		return Fail, ""
	}
	defer file.Close()

	helper.LogInfo(fmt.Sprintf("========= Starting to assemble '%s' =========", path))

	ptrProgram := programFromFile(file)
	if ptrProgram == nil {
		return Fail, ""
	}

	errs := ptrProgram.ReplaceLabelsWithNumbers()
	if len(errs) > 0 {
		for _, err := range errs {
			// TODO: create infrastructure go get lineIndex and line here
			helper.LogErrorInLine(*myerrors.NewCodeError(err), 0, "")
		}
		return Fail, ""
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
		return Fail, ""
	}

	resultCode := writeExecProgram(*ptrProgram, execFileName, execFile)
	if resultCode != 0 {
		return Fail, ""
	}

	return Success, execFileName
}

func programFromFile(file io.Reader) *data.Program {
	reader := bufio.NewReader(file)
	lineIndex := 1
	program := data.NewProgram(0)

	successful := true

	for {
		line, errRead := reader.ReadString('\n')

		if errRead != nil && errRead != io.EOF {
			helper.LogOtherError(fmt.Sprintf("Error while reading file. Error: %s", errRead.Error()))
			return nil
		}

		jumpLabel, commandPointer, errs := assembleEntireLine(line)

		if jumpLabel != nil {
			program.AddJumpLabel(*jumpLabel, program.LenCommands())
		}

		if len(errs) > 0 {
			successful = false

			for _, err := range errs {
				helper.LogErrorInLine(err, lineIndex, line)
			}
		} else if commandPointer != nil {
			program.AddCommand(*commandPointer)
		}

		if errRead == io.EOF {
			break
		}
		lineIndex++
	}

	if !successful {
		return nil
	}
	return &program
}

func assembleEntireLine(line string) (*string, *data.Command, []myerrors.CustomError) {
	normalizedStr := utils.LineNormalization(line)

	if normalizedStr == "" {
		return nil, nil, nil
	}

	errs := make([]myerrors.CustomError, 0)

	jumpLabel, restOfCommandStr, errLabel := AssembleJumpLabel(normalizedStr)
	if errLabel != nil {
		errs = append(errs, *errLabel)
	}

	if restOfCommandStr == "" {
		return jumpLabel, nil, errs
	}

	commandPointer, errCmd := AssembleCommand(restOfCommandStr)

	if errCmd != nil {
		errs = append(errs, *errCmd)
	}

	return jumpLabel, commandPointer, errs
}

func writeExecProgram(program data.Program, execFileName string, execFile io.Writer) int {
	writer := bufio.NewWriter(execFile)
	defer writer.Flush()

	execStr, errs := program.ToExecuter()

	if errs != nil && len(errs) > 0 {
		for _, err := range errs {
			// TODO: infrastructure to get line
			helper.LogErrorInLine(err, 0, "")
		}
		return 1
	}

	_, err := writer.WriteString(execStr)
	if err != nil {
		helper.LogOtherError(fmt.Sprintf("Could not write to file %s \n", execFileName))
		return 2
	}

	return 0
}
