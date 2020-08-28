package instructionsexplanation

import "testing"

func TestAmountInstructions(t *testing.T) {
	cmdExplanations := GetInstructionsExplanation()

	for name, explanation := range cmdExplanations {
		if explanation.Instruction == "" || explanation.Param == "" {
			t.Errorf("%s explanation is not complete", name)
		}
	}
}

// TODO:
// func GetInstructionsExplanationSorted() ([]string, map[string]InstructionExplanation) {
// func GetInstructionExplanation(name string) (*InstructionExplanation, error) {
