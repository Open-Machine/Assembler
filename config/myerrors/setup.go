package myerrors

import (
	"fmt"
)

func CommandAlreadyExistError(instructionStr string) error {
	return fmt.Errorf("Map key '%s' already exists, there can't be any commands with the same name", instructionStr)
}
