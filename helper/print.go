package helper

import (
	"assembler/myerrors"
	"fmt"
	"io"
	"os"
)

var Out io.Writer = os.Stdout
var Err io.Writer = os.Stderr

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorPurple = "\033[35m"
var colorYellow = "\033[33m"
var colorWhite = "\033[37m"

func LogInfo(str string) {
	formatedStr := fmt.Sprintf("[INFO] %s \n", str)
	println(Out, colorWhite, formatedStr)
}

func PrintlnExplanation(str string) {
	println(Out, colorWhite, str)
}

func LogWarning(str string) {
	formatedStr := fmt.Sprintf("[Warning] %s \n", str)
	println(Err, colorYellow, formatedStr)
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
		lastLine := fmt.Sprintf("Please consider openning an ISSUE on Github: https://github.com/Open-Machine/Assembler/issues .")

		beginningLines := append([]string{firstLine}, middleLines...)
		lines = append(beginningLines, lastLine)
	} else {
		color = colorRed

		firstLine := fmt.Sprintf("[ERROR] Error on line %d", lineIndex)
		lines = append([]string{firstLine}, middleLines...)
	}

	for _, formatedStr := range lines {
		println(Err, string(color), formatedStr)
	}
}

func LogOtherError(str string) {
	formatedStr := fmt.Sprintf("[ERROR] %s", str)
	println(Err, colorRed, formatedStr)
}

func println(writer io.Writer, color string, str string) {
	fmt.Fprintln(writer, string(color), str, string(colorReset))
}
