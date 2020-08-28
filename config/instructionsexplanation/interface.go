package instructionsexplanation

import (
	"sort"

	"github.com/open-machine/assembler/config/myerrors"
)

func GetInstructionsExplanation() map[string]InstructionExplanation {
	return instructionExplanations
}

func GetInstructionsExplanationSorted() ([]string, map[string]InstructionExplanation) {
	keys := make([]string, len(instructionExplanations))
	i := 0
	for k := range instructionExplanations {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	return keys, instructionExplanations
}

func GetInstructionExplanation(name string) (*InstructionExplanation, error) {
	explanation, exists := instructionExplanations[name]
	if !exists {
		return nil, myerrors.InstructionDoesNotExistError(name)
	}
	return &explanation, nil
}
