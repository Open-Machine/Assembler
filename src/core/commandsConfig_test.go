package core

import "testing"

func Test(t *testing.T) {
	cmdExplanations := GetCommandsExplanation()

	if len(cmdExplanations) != 12 {
		t.Errorf("Expected 12 commands explanations")
	}

	for name, explanation := range cmdExplanations {
		if explanation.Command == "" || explanation.Param == "" {
			t.Errorf("%s explanation is not complete", name)
		}
	}
}
