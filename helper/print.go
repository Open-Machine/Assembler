package helper

import (
	"fmt"
	"io"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/myerrors"
)

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorPurple = "\033[35m"
var colorYellow = "\033[33m"
var colorWhite = "\033[37m"
var _ = "\033[32m" //colorGreen

func LogStep(str string) {
	formatedStr := fmt.Sprintf("========= %s =========", str)
	LogInfo(formatedStr)
}

func LogInfo(str string) {
	formatedStr := fmt.Sprintf("[INFO] %s \n", str)
	println(config.Out, colorWhite, formatedStr)
}

func PrintlnExplanation(str string) {
	println(config.Out, colorWhite, str)
}

func LogWarning(str string) {
	formatedStr := fmt.Sprintf("[Warning] %s \n", str)
	println(config.Err, colorYellow, formatedStr)
}

func LogErrorInLine(customErr myerrors.CustomError, lineIndex int, line string) {
	middleLines := []string{
		fmt.Sprintf("\t\tError: %s", customErr.Error()),
		fmt.Sprintf("\t\tLine: %s", line),
	}

	var color string
	var lines []string
	if !customErr.IsCodeError() {
		color = colorPurple

		firstLine := fmt.Sprintf("[ASSEMBLER ERROR] Unexpected error while compiling line %d.", lineIndex)
		lastLine := "Please consider openning an ISSUE on Github: https://github.com/Open-Machine/github.com/open-machine/assembler/issues ."

		beginningLines := append([]string{firstLine}, middleLines...)
		lines = append(beginningLines, lastLine)
	} else {
		color = colorRed

		firstLine := fmt.Sprintf("[ERROR] Error on line %d", lineIndex)
		lines = append([]string{firstLine}, middleLines...)
	}

	for _, formatedStr := range lines {
		println(config.Err, string(color), formatedStr)
	}
}

func LogOtherError(str string) {
	formatedStr := fmt.Sprintf("[ERROR] %s", str)
	println(config.Err, colorRed, formatedStr)
}

func println(writer io.Writer, color string, str string) {
	fmt.Fprintln(writer, string(color), str, string(colorReset))
}
