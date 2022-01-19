package myerrors

import (
	"fmt"
)

func MoreThanOneCommandInLine(restOfLine string) error {
	return fmt.Errorf("each line can only have one command. It cannot be more than one of the following: a jump label, and an instruction. '%s' should be in the next line", restOfLine)
}

func WrongAssemblerState(gotLine string) error {
	return fmt.Errorf("Wrong assembler state: did not expect '%s'", gotLine)
}
