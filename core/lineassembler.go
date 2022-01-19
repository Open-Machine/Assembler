package core

import (
	"strings"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/myerrors"
	"github.com/open-machine/assembler/core/comment"
	"github.com/open-machine/assembler/core/instruction"
	"github.com/open-machine/assembler/core/label"
	"github.com/open-machine/assembler/core/variable"
	"github.com/open-machine/assembler/data"
	"github.com/open-machine/assembler/utils"
)

func getLineWithoutComment(line string) string {
	normalizedStr := utils.LineNormalization(line)
	lineWithoutComment := comment.RemoveComment(normalizedStr)
	return strings.TrimSpace(lineWithoutComment)
}

func nextAssemblerState(lineWithoutComment string, currentState int) (int, *myerrors.CustomError) {
	switch lineWithoutComment {
	case config.AssemblerStateVarHeader:
		if currentState == NoneState {
			return VarsState, nil
		} else {
			return -1, myerrors.NewCodeError(myerrors.WrongAssemblerState(lineWithoutComment))
		}
	case config.AssemblerStateCodeHeader:
		if currentState == CodeState {
			return -1, myerrors.NewCodeError(myerrors.WrongAssemblerState(lineWithoutComment))
		} else {
			// it can be VarsState or NoneState, because it doens't need to have @VARS defined first
			return CodeState, nil
		}
	}
	return -1, nil // it doesn't change the state
}

func assembleEntireLine(line string, assemblerState *int, variableIndex *int) (*string, *data.Variable, *data.Instruction, *myerrors.CustomError) {
	lineWithoutComment := getLineWithoutComment(line)
	if lineWithoutComment == "" {
		return nil, nil, nil, nil
	}

	updatedState, errState := nextAssemblerState(lineWithoutComment, *assemblerState)
	if errState != nil {
		return nil, nil, nil, errState
	}
	if updatedState >= 0 {
		*assemblerState = updatedState
		return nil, nil, nil, nil
	} else {
		switch *assemblerState {
		case VarsState:
			variable, varErr := variable.AssembleVariable(lineWithoutComment, variableIndex)
			if varErr != nil {
				return nil, nil, nil, varErr
			}
			return nil, variable, nil, nil
		case CodeState:
			// Jump Label
			jumpLabel, restOfLine, errLabel := label.AssembleJumpLabel(lineWithoutComment)
			if errLabel != nil {
				return nil, nil, nil, errLabel
			}
			if jumpLabel != nil && restOfLine != "" {
				return nil, nil, nil, myerrors.NewCodeError(myerrors.MoreThanOneCommandInLine(restOfLine))
			}
			if jumpLabel != nil {
				return jumpLabel, nil, nil, nil
			}

			// Instruction
			instructionPointer, errInstruc := instruction.AssembleInstruction(restOfLine)
			if errInstruc != nil {
				return nil, nil, nil, errInstruc
			}

			return jumpLabel, nil, instructionPointer, nil
		case NoneState:
		}
		return nil, nil, nil, myerrors.NewCodeError(myerrors.WrongAssemblerState(lineWithoutComment))
	}
}
