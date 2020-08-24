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
		helper.PrintlnErr(fmt.Sprintf("Cannot open file %s . Error: %s", path, err.Error()))
		return
	}
	defer file.Close()

	helper.PrintlnOut(fmt.Sprintf("========= Starting to assemble %s =========", path))

	ptrProgram := programFromFile(file)
	if ptrProgram == nil {
		printFailedToAssemble(path)
		return
	}

	errs := ptrProgram.ReplaceLabelsWithNumbers()
	if len(errs) > 0 {
		for _, err := range errs {
			// TODO: create infrastructure go get lineIndex and line here
			printErrorOnLine(0, "", *myerrors.NewCodeError(err))
		}
		printFailedToAssemble(path)
		return
	}

	binaryFileName := helper.FilenameWithoutExtension(file.Name())

	binaryFile, err := os.Create(binaryFileName)
	if err != nil {
		helper.PrintlnErr(fmt.Sprintf("Cannot open file %s . Error: %s", binaryFileName, err.Error()))
		printFailedToAssemble(path)
		return
	}

	resultCode := writeBinaryProgram(*ptrProgram, binaryFileName, binaryFile)
	if resultCode != 0 {
		printFailedToAssemble(path)
		return
	}
	helper.PrintlnOut(fmt.Sprintf("========= Binary file available in %s =========", binaryFileName))
}

func programFromFile(file io.Reader) *data.Program {
	reader := bufio.NewReader(file)
	lineIndex := 1
	program := data.NewProgram(0)

	successful := true

	for {
		line, errRead := reader.ReadString('\n')

		if errRead != nil && errRead != io.EOF {
			helper.PrintlnErr(fmt.Sprintf("Error while reading file. Error: %s", errRead.Error()))
			return nil
		}

		gotoLabel, commandPointer, errs := assembleEntireLine(line)

		if gotoLabel != nil {
			program.AddGotoLabel(*gotoLabel, program.LenCommands())
		}

		if len(errs) > 0 {
			successful = false

			for _, err := range errs {
				printErrorOnLine(lineIndex, line, err)
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

	gotoLabel, restOfCommandStr, errLabel := AssembleGotoLabel(normalizedStr)
	if errLabel != nil {
		errs = append(errs, *errLabel)
	}

	if restOfCommandStr == "" {
		return gotoLabel, nil, errs
	}

	commandPointer, errCmd := AssembleCommand(restOfCommandStr)

	if errCmd != nil {
		errs = append(errs, *errCmd)
	}

	return gotoLabel, commandPointer, errs
}

func writeBinaryProgram(program data.Program, binaryFileName string, binaryFile io.Writer) int {
	writer := bufio.NewWriter(binaryFile)
	defer writer.Flush()

	binaryStr, errs := program.ToExecuter()

	if errs != nil && len(errs) > 0 {
		helper.PrintlnErr("[Assembler Errors] It was not possible to assemble your code. Encountered errors:")
		for i, err := range errs {
			helper.PrintlnErr(fmt.Sprintf("\t[%d] %s", i+1, err.Error()))
		}
		return 1
	}

	_, err := writer.WriteString(binaryStr)
	if err != nil {
		helper.PrintlnErr(fmt.Sprintf("Could not write to file %s \n", binaryFileName))
		return 2
	}

	return 0
}

func printFailedToAssemble(path string) {
	helper.PrintlnErr(fmt.Sprintf("========= Failed to assemble %s =========", path))
}

func printErrorOnLine(lineIndex int, line string, err myerrors.CustomError) {
	line = utils.RemoveNewLine(line)

	helper.PrintlnErr(fmt.Sprintf("[Error] Error on line %d: '%s'", lineIndex, line))
	helper.PrintlnErr(fmt.Sprintf("\t\tError: %s", err.Error()))
	helper.PrintlnErr("")
}
