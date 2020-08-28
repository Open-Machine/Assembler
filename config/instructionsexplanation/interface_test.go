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

func TestOrderExplanations(t *testing.T) {
	var tests = []struct {
		index    int
		expected string
	}{
		{0, "add"},
		{1, "copy"},
	}

	for _, test := range tests {
		keys, _ := GetInstructionsExplanationSorted()

		if keys[test.index] != test.expected {
			t.Errorf("Explanation order is wrong")
		}
	}
}

func TestSpecificExplanation(t *testing.T) {
	var tests = []struct {
		param  string
		exists bool
	}{
		{"add", true},
		{"copy", true},
		{"mov", false},
	}

	for _, test := range tests {
		_, err := GetInstructionExplanation(test.param)
		gotErr := err != nil

		if !gotErr != test.exists {
			t.Errorf("Explanation missing or one more")
		}
	}
}
