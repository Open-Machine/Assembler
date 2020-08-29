package myerrors

import (
	"errors"
	"fmt"
)

func MoreThanOneCommandInLine(restOfLine string) error {
	return fmt.Errorf("Each line can only have one command. It cannot be more than one of the following: a jump label, and an instruction. '%s' should be in the next line.", restOfLine)
}

func InvalidStateTransformationToExecuterError() error {
	return errors.New("Invalid State: Cannot transform instruction to executer while parameter is still a label")
}
