// This programs reads the outputs of logisim-evolution and formats to readable lines.
// This script requires a single argument that represents the print format of the numbers, the options are:
// - b: for binary
// - h: for hexadecimal
// - d: for decimal

package main

import (
	"fmt"
	"os"
	"strconv"
)

func binaryToInt(binaryStr string) (int64, error) {
	formattedStr := binaryStr[1:]
	if binaryStr[0] == '1' {
		formattedStr = "-" + formattedStr
	}
	return strconv.ParseInt(formattedStr, 2, 16)
}

func binaryToHexadecimal(num int) string {
	return fmt.Sprintf("%x", num)
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: this script requires a single argument that represents the print format of the numbers")
		return
	}
	numFormat := os.Args[1]

	var str1, str2, str3, str4 string
	_, err := fmt.Scanf("%s %s %s %s", &str1, &str2, &str3, &str4)
	binaryStr := str1 + str2 + str3 + str4
	if err != nil {
		panic(err)
	}
	fmt.Println(format(binaryStr, numFormat))
}
