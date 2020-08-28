package instruction

import "testing"

func TestAmountInstructions(t *testing.T) {
	cmdExplanations := GetInstructionsExplanation()

	if len(cmdExplanations) != 12 {
		t.Errorf("Expected 12 instructions explanations")
	}

	for name, explanation := range cmdExplanations {
		if explanation.Instruction == "" || explanation.Param == "" {
			t.Errorf("%s explanation is not complete", name)
		}
	}
}
