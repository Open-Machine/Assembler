package myerrors

import (
	"fmt"
)

func InstructionCodeOverflow(instructionCode int, amntBits int) error {
	return fmt.Errorf("Instruction code '%b' overflows %d bits", instructionCode, amntBits)
}

func InstructionDoesNotExistError(instructionStr string) error {
	return fmt.Errorf("Instruction '%s' does not exist", instructionStr)
}
