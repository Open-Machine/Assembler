package instruction

import (
	"testing"

	"github.com/open-machine/assembler/config/instructionsexplanation"
)

func TestAmountInstructions(t *testing.T) {
	if len(instructions) != 12 {
		t.Errorf("Expected 12 instructions explanations! Change 'intructions_test.go' and 'config/instructionsexplanation/*'")
	}
}

func TestSameKeysAsConfig(t *testing.T) {
	instructionsExplanation := instructionsexplanation.GetInstructionsExplanation()

	if len(instructions) != len(instructionsExplanation) {
		t.Errorf("Amnt of explanations on the config and explanations are different!")
	}

	for explanationKey, _ := range instructionsExplanation {
		_, exists := instructions[explanationKey]
		if !exists {
			t.Errorf("Instruction '%s' has explanation but no configuration!", explanationKey)
		}
	}

	for configKey, _ := range instructions {
		_, exists := instructionsExplanation[configKey]
		if !exists {
			t.Errorf("Instruction '%s' has configuration but no explanation!", configKey)
		}
	}

}
