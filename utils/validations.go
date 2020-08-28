package utils

import (
	"math"
	"regexp"

	"github.com/open-machine/assembler/config"
)

func IsOverflow(num uint, availableBits int) bool {
	largestNumber := math.Pow(2, float64(availableBits))
	return num >= uint(math.Floor(largestNumber))
}

func IsValidVarName(str string) bool {
	matched, err := regexp.MatchString(config.VariableNameRegex, str)
	return matched && err == nil
}

// TODO: check valid name
// TODO: use to validate name
// func isReservedWord(str string) bool {
// }
