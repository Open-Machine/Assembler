package utils

import (
	"strconv"

	"github.com/open-machine/assembler/config/myerrors"
)

const bitsize = 64 // so that numStr is never considered negative
func StrToPositiveInt(numStr string) (uint, error) {
	isHex := len(numStr) >= 2 && numStr[0:2] == "0x"
	if isHex {
		hexStr := numStr[2:]

		num, err := strconv.ParseUint(hexStr, 16, bitsize)
		return uint(num), err
	}

	num, err := strconv.ParseInt(numStr, 10, bitsize)
	return uint(num), err
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
