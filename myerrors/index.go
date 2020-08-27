package myerrors

import (
	"errors"
	"fmt"
)

func CommandCodeOverflow(commandCode int, amntBits int) error {
	return fmt.Errorf("Command code '%b' overflows %d bits", commandCode, amntBits)
}

func ParamOverflow(param int, amntBits int) error {
	return fmt.Errorf("Param '%b' overflows %d bits", param, amntBits)
}

func JumpLabelAlreadyExistsError(label string) error {
	return fmt.Errorf("Jump label '%s' already exists", label)
}

func JumpLabelDoesNotExistError(label string) error {
	return fmt.Errorf("Jump label '%s' does not exist", label)
}

func InvalidLabelParam(label string) error {
	return fmt.Errorf("Param '%s' is not a valid label name", label)
}

func InvalidParamLabelOrInt(param string, err error) error {
	return fmt.Errorf("Param '%s' is not a valid label nor a valid number (Conversion error: %s)", param, err.Error())
}

func InvalidParamInt(param string, err error) error {
	return fmt.Errorf("Param '%s' is not a valid number (Conversion error: %s)", param, err.Error())
}

func InvalidStateTransformationToExecuterError() error {
	return errors.New("Invalid State: Cannot transform command to executer while parameter is still a label")
}

func WrongNumberOfParamsError(command string, amntExpected int, amntReceived int, params []string) error {
	strParameters := ""
	if len(params) == 0 {
		strParameters = "no params"
	} else {
		for i, param := range params {
			strParameters += fmt.Sprintf("'%s'", param)
			if i != len(params)-1 {
				strParameters += ", "
			}
		}
	}

	return fmt.Errorf("The command '%s' requires %d parameters, but received %d parameters (parameters: %s)", command, amntExpected, amntReceived, strParameters)
}

func CommandDoesNotExistError(commandStr string) error {
	return fmt.Errorf("Command '%s' does not exist", commandStr)
}

func InvalidNumberParamParseToHexStrError(num int, strLength int, hexStr string) error {
	return fmt.Errorf("Number %d cannot be converted to hexadecimal string of length %d. Got: '%s'", num, strLength, hexStr)
}
