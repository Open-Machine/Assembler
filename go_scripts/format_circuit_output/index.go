package main

// This programs reads the outputs of logisim-evolution and formats to readable lines.
// This script requires a single argument that represents the print format of the numbers, the options are:
// - b: for binary
// - h: for hexadecimal
// - d: for decimal

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func binaryToInt(binaryStr string) (int64, error) {
	formattedStr := binaryStr[1:]
	if binaryStr[0] == '1' {
		formattedStr = "-" + formattedStr
	}
	return strconv.ParseInt(formattedStr, 2, 16)
}

func binaryToHexadecimal(num int) string {
	isNegative := num < 0
	abs := num
	preffix := ""
	if isNegative {
		abs = -num
		preffix = "-"
	}

	formatted := fmt.Sprintf("%x", abs)
	return preffix + "0x" + strings.ToUpper(formatted)
}
func binaryToDecimal(num int) string {
	return fmt.Sprintf("%d", num)
}

func format(binaryStr string, numFormat string) string {
	switch numFormat {
	case "b":
		return binaryStr
	case "h", "d":
		num, err := binaryToInt(binaryStr)
		if err != nil {
			return fmt.Sprintf("erro: output desconhecido (recebeu: %s)", binaryStr)
		}
		switch numFormat {
		case "h":
			return binaryToHexadecimal(int(num))
		case "d":
			return binaryToDecimal(int(num))
		default:
			return "error"
		}
	default:
		return "error: unknown format"
	}
}

func program(args []string, consoleReader io.Reader, consoleWriter io.Writer) {
	if len(args) < 2 {
		fmt.Fprintln(consoleWriter, "error: this script requires a single argument that represents the print format of the numbers")
		return
	}
	numFormat := args[1]

	isFirst := true // first print is always zero
	for {
		var str1, str2, str3, str4, str5 string
		_, err := fmt.Fscanf(consoleReader, "%s %s %s %s     %s", &str1, &str2, &str3, &str4, &str5)
		if isFirst {
			isFirst = false
			continue
		}

		binaryStr := str1 + str2 + str3 + str4
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(consoleWriter, format(binaryStr, numFormat))

	}
}

func main() {
	program(os.Args, os.Stdin, os.Stdout)
}
