package instruction

import "testing"

func TestAmountInstructions(t *testing.T) {
	if len(instructions) != 12 {
		t.Errorf("Expected 12 instructions explanations! ALSO CHANGE 'INSTRUCTION_TEST.GO'")
	}
}
