package label

import (
	"strings"

	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/utils"
)

func AssembleJumpLabel(line string) (*string, string, *myerrors.CustomError) {
	indexOfColon := strings.Index(line, ":")

	if indexOfColon < 0 {
		return nil, line, nil
	}

	label := strings.TrimSpace(line[:indexOfColon])
	restOfLine := line[indexOfColon+1:]

	err := utils.CheckJumpLabelName(label)

	if err != nil {
		return nil, restOfLine, myerrors.NewCodeError(err)
	}

	return &label, restOfLine, nil
}
