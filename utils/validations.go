package utils

import (
	"math"
	"regexp"

	"github.com/open-machine/assembler/config/myerrors"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/instructionsexplanation"
)

func IsOverflow(num uint, availableBits int) bool {
	largestNumber := math.Pow(2, float64(availableBits))
	return num >= uint(math.Floor(largestNumber))
}

func IsValidName(varName string) bool {
	return nameStatus(varName) == validName
}

const (
	validName        = iota
	invalidNameRegex = iota
	reservedWord     = iota
)

func nameStatus(varName string) int {
	if isReservedWord(varName) {
		return reservedWord
	}

	matched, err := regexp.MatchString(config.VariableNameRegex, varName)
	isValidName := matched && err == nil
	if !isValidName {
		return invalidNameRegex
	}
	return validName
}

func isReservedWord(varName string) bool {
	instructions := instructionsexplanation.GetInstructionsExplanation()
	for instructionName, _ := range instructions {
		if varName == instructionName {
			return true
		}
	}
	return false
}

func CheckParamName(varName string) error {
	status := nameStatus(varName)
	if status == invalidNameRegex {
		return myerrors.InvalidParamLabel(varName)
	}
	if status == reservedWord {
		return myerrors.ReservedWordParamError(varName)
	}
	return nil
}

func CheckJumpLabelName(varName string) error {
	status := nameStatus(varName)
	if status == invalidNameRegex {
		return myerrors.InvalidJumpLabel(varName)
	}
	if status == reservedWord {
		return myerrors.ReservedWordJumpLabelError(varName)
	}
	return nil
}
