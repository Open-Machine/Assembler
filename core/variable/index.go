package variable

import (
	"fmt"
	"strings"

	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/utils"
)

func AssembleVariable(line string, varIndex *int) (*data.Variable, *myerrors.CustomError) {
	equalIndex := strings.Index(line, "=")
	if equalIndex < 0 {
		return nil, myerrors.NewAssemblerError(fmt.Errorf("'=' was expected"))
	}

	variableName := strings.TrimSpace(line[:equalIndex])

	initialValueStr := strings.TrimSpace(line[equalIndex+1:])
	initialValue, err2 := utils.StrToPositiveInt(initialValueStr)
	if err2 != nil {
		return nil, myerrors.NewAssemblerError(err2)
	}

	variable, err := data.NewVariable(variableName, *varIndex, initialValue)
	if err != nil {
		return nil, err
	}

	*varIndex++
	return variable, nil
}
