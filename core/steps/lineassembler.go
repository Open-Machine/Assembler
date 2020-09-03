package steps

import (
	"github.com/open-machine/assembler/commands/instruction"
	"github.com/open-machine/assembler/commands/label"
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/utils"
)

func assembleEntireLine(line string) (*string, *data.Instruction, *myerrors.CustomError) {
	normalizedStr := utils.LineNormalization(line)

	if normalizedStr == "" {
		return nil, nil, nil
	}

	// Jump Label
	jumpLabel, restOfLine, errLabel := label.AssembleJumpLabel(normalizedStr)
	if errLabel != nil {
		return nil, nil, errLabel
	}
	if jumpLabel != nil {
		if restOfLine != "" {
			return nil, nil, myerrors.NewCodeError(myerrors.MoreThanOneCommandInLine(restOfLine))
		}
		return jumpLabel, nil, nil
	}

	// Instruction
	instructionPointer, errInstruc := instruction.AssembleInstruction(restOfLine)
	if errInstruc != nil {
		return nil, nil, errInstruc
	}

	return jumpLabel, instructionPointer, nil
}
