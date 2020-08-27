package core

import (
	"assembler/data"
	"assembler/helper"
	"assembler/myerrors"
	"assembler/utils"
	"bufio"
	"fmt"
	"io"
	"os"
)

func AssembleFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		helper.LogOtherError(fmt.Sprintf("Cannot open file %s . Error: %s", path, err.Error()))
		return
	}
	defer file.Close()

	helper.LogInfo(fmt.Sprintf("========= Starting to assemble %s =========", path))

	ptrProgram := programFromFile(file)
	if ptrProgram == nil {
		printFailedToAssemble(path)
		return
	}

	errs := ptrProgram.ReplaceLabelsWithNumbers()
	if len(errs) > 0 {
		for _, err := range errs {
			// TODO: create infrastructure go get lineIndex and line here
			helper.LogErrorInLine(*myerrors.NewCodeError(err), 0, "")
		}
		printFailedToAssemble(path)
		return
	}

	binaryFileName := helper.FilenameWithoutExtension(file.Name())

	binaryFile, err := os.Create(binaryFileName)
	if err != nil {
		helper.LogOtherError(fmt.Sprintf("Cannot open file %s . Error: %s", binaryFileName, err.Error()))
		printFailedToAssemble(path)
		return
	}

	resultCode := writeBinaryProgram(*ptrProgram, binaryFileName, binaryFile)
	if resultCode != 0 {
		printFailedToAssemble(path)
		return
	}
	helper.LogInfo(fmt.Sprintf("========= Binary file available in %s =========", binaryFileName))
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

func writeBinaryProgram(program data.Program, binaryFileName string, binaryFile io.Writer) int {
	writer := bufio.NewWriter(binaryFile)
	defer writer.Flush()

	binaryStr, errs := program.ToExecuter()

	if errs != nil && len(errs) > 0 {
		for _, err := range errs {
			// TODO: infrastructure to get line
			helper.LogErrorInLine(err, 0, "")
		}
		return 1
	}

	_, err := writer.WriteString(binaryStr)
	if err != nil {
		helper.LogOtherError(fmt.Sprintf("Could not write to file %s \n", binaryFileName))
		return 2
	}

	return 0
}

func printFailedToAssemble(path string) {
	helper.LogOtherError(fmt.Sprintf("========= Failed to assemble %s =========", path))
}
