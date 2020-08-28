package core

import (
	"github.com/open-machine/assembler/core/instruction"
	"github.com/open-machine/assembler/core/label"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/myerrors"
	"github.com/open-machine/assembler/utils"
)

func assembleEntireLine(line string) (*string, *data.Instruction, []myerrors.CustomError) {
	normalizedStr := utils.LineNormalization(line)

	if normalizedStr == "" {
		return nil, nil, nil
	}

	errs := make([]myerrors.CustomError, 0)

	jumpLabel, restOfInstructionStr, errLabel := label.AssembleJumpLabel(normalizedStr)
	if errLabel != nil {
		errs = append(errs, *errLabel)
	}

	if restOfInstructionStr == "" {
		return jumpLabel, nil, errs
	}

	instructionPointer, errCmd := instruction.AssembleInstruction(restOfInstructionStr)

	if errCmd != nil {
		errs = append(errs, *errCmd)
	}

	return jumpLabel, instructionPointer, errs
}
