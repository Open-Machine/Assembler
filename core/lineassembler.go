package core

import (
	"strings"

	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/core/comment"
	"github.com/open-machine/assembler/core/instruction"
	"github.com/open-machine/assembler/core/label"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/utils"
)

func assembleEntireLine(line string) (*string, *data.Instruction, *myerrors.CustomError) {
	normalizedStr := utils.LineNormalization(line)

	lineWithoutCommand := comment.RemoveComment(normalizedStr)
	lineWithoutCommand = strings.TrimSpace(lineWithoutCommand)

	if lineWithoutCommand == "" {
		return nil, nil, nil
	}

	// Jump Label
	jumpLabel, restOfLine, errLabel := label.AssembleJumpLabel(lineWithoutCommand)
	if errLabel != nil {
		return nil, nil, errLabel
	}
	if jumpLabel != nil && restOfLine != "" {
		return nil, nil, myerrors.NewCodeError(myerrors.MoreThanOneCommandInLine(restOfLine))
	}
	if jumpLabel != nil {
		return jumpLabel, nil, nil
	}

	// Instruction
	instructionPointer, errInstruc := instruction.AssembleInstruction(restOfLine)
	if errInstruc != nil {
		return nil, nil, errInstruc
	}

	return jumpLabel, instructionPointer, nil
}
