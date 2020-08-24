package utils

import (
	"assembler/config"
	"math"
	"regexp"
)

func IsOverflow(num uint, availableBits int) bool {
	largestNumber := math.Pow(2, float64(availableBits))
	return num >= uint(math.Floor(largestNumber))
}

func IsValidVarName(str string) bool {
	matched, err := regexp.MatchString(config.VariableNameRegex, str)
	return matched && err == nil
}
