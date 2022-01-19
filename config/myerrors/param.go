package myerrors

import "fmt"

func ValueOverflow(number uint, amntBits int) error {
	return fmt.Errorf("Value '%b' overflows %d bits", number, amntBits)
}

func InvalidParamLabelOrInt(param string, err error) error {
	return fmt.Errorf("Param '%s' is not a valid label nor a valid number (Conversion error: %s)", param, err.Error())
}

func InvalidParamInt(param string, err error) error {
	return fmt.Errorf("Param '%s' is not a valid number (Conversion error: %s)", param, err.Error())
}

func WrongNumberOfParamsError(instruction string, amntExpected int, amntReceived int, params []string) error {
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

	return fmt.Errorf("The instruction '%s' requires %d parameters, but received %d parameters (parameters: %s)", instruction, amntExpected, amntReceived, strParameters)
}

func InvalidNumberParamParseToHexStrError(num int, strLength int, hexStr string) error {
	return fmt.Errorf("Number %d cannot be converted to hexadecimal string of length %d. Got: '%s'", num, strLength, hexStr)
}
