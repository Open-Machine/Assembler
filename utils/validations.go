package utils

import (
	"math"
	"regexp"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/instructionsexplanation"
)

func IsOverflow(num uint, availableBits int) bool {
	largestNumber := math.Pow(2, float64(availableBits))
	return num >= uint(math.Floor(largestNumber))
}

func IsValidVarName(varName string) bool {
	matched, err := regexp.MatchString(config.VariableNameRegex, varName)
	return matched && err == nil
}

// TODO: check valid name
// TODO: use to validate name
func isReservedWord(varName string) bool {
	instructions := instructionsexplanation.GetInstructionsExplanation()
	for instructionName, _ := range instructions {
		if varName == instructionName {
			return true
		}
	}
	return false
}
