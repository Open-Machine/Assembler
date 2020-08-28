package core

import (
	"github.com/open-machine/assembler/myerrors"
	"github.com/open-machine/assembler/utils"
	"strings"
)

func AssembleJumpLabel(line string) (*string, string, *myerrors.CustomError) {
	indexOfColon := strings.Index(line, ":")

	if indexOfColon < 0 {
		return nil, line, nil
	}

	label := strings.TrimSpace(line[:indexOfColon])
	restOfLine := strings.TrimSpace(line[indexOfColon+1:])

	if !utils.IsValidVarName(label) {
		return nil, restOfLine, myerrors.NewCodeError(myerrors.InvalidLabelParam(label))
	}

	return &label, restOfLine, nil
}
