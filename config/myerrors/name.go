package myerrors

import (
	"fmt"

	"github.com/open-machine/assembler/config"
)

func JumpLabelAlreadyExistsError(label string) error {
	return fmt.Errorf("Jump label '%s' already exists", label)
}

func JumpLabelDoesNotExistError(label string) error {
	return fmt.Errorf("Jump label '%s' does not exist", label)
}

// Reserved Word

func reservedWordError(label string, commandType string) error {
	return fmt.Errorf("'%s' is reserved word, it cannot be used as a %s", label, commandType)
}

func ReservedWordParamError(label string) error {
	return reservedWordError(label, "param")
}

func ReservedWordJumpLabelError(label string) error {
	return reservedWordError(label, "jump label")
}

// Invalid Name

func invalidName(label string, first string) error {
	return fmt.Errorf("%s '%s' is not a valid name, it should follow the following regex: '%s'", first, label, config.VariableNameRegex)
}

func InvalidParamLabel(label string) error {
	return invalidName(label, "Param")
}

func InvalidJumpLabel(label string) error {
	return invalidName(label, "Jump label")
}
