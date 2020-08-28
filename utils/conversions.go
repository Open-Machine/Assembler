package utils

import (
	"github.com/open-machine/assembler/myerrors"
	"strconv"
)

const bitsize = 64 // so that numStr is never considered negative
func StrToPositiveInt(numStr string) (int, error) {
	isHex := len(numStr) >= 2 && numStr[0:2] == "0x"
	if isHex {
		hexStr := numStr[2:]

		num, err := strconv.ParseUint(hexStr, 16, bitsize)
		return int(num), err
	}

	num, err := strconv.ParseInt(numStr, 10, bitsize)
	return int(num), err
}

func IntToStrHex(num int, strLength int) (string, error) {
	hexStr := strconv.FormatInt(int64(num), 16)

	if len(hexStr) > strLength {
		return "", myerrors.InvalidNumberParamParseToHexStrError(num, strLength, hexStr)
	}

	for len(hexStr) < strLength {
		hexStr = "0" + hexStr
	}

	return hexStr, nil
}
