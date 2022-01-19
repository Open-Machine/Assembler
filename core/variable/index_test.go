package variable

import (
	"testing"

	"github.com/open-machine/assembler/data"
)

func TestAssembleVariable(t *testing.T) {
	var tests = []struct {
		line                string
		expectedVariable    *data.Variable
		shouldIncreaseIndex bool
		expectsErr          bool
	}{
		// spaced
		{
			line:                "var = 1",
			expectedVariable:    newVariable("var", 0, 1),
			shouldIncreaseIndex: true,
			expectsErr:          false,
		},
		// too spaced
		{
			line:                "     var    =    1  ",
			expectedVariable:    newVariable("var", 0, 1),
			shouldIncreaseIndex: true,
			expectsErr:          false,
		},
		// zero space
		{
			line:                "var=1",
			expectedVariable:    newVariable("var", 0, 1),
			shouldIncreaseIndex: true,
			expectsErr:          false,
		},
		// different numbers
		{
			line:                "var=0x1",
			expectedVariable:    newVariable("var", 0, 1),
			shouldIncreaseIndex: true,
			expectsErr:          false,
		},
		// number overflow
		{
			line:                "var = 200000",
			expectedVariable:    nil,
			shouldIncreaseIndex: false,
			expectsErr:          true,
		},
		// no equal
		{
			line:                "var  0x1",
			expectedVariable:    nil,
			shouldIncreaseIndex: false,
			expectsErr:          true,
		},
		// no number
		{
			line:                "var",
			expectedVariable:    nil,
			shouldIncreaseIndex: false,
			expectsErr:          true,
		},
		{
			line:                "var =",
			expectedVariable:    nil,
			shouldIncreaseIndex: false,
			expectsErr:          true,
		},
		{
			line:                "var = true",
			expectedVariable:    nil,
			shouldIncreaseIndex: false,
			expectsErr:          true,
		},
	}

	for i, test := range tests {
		index := 0
		got, err := AssembleVariable(test.line, &index)

		if (got == nil) != (test.expectedVariable == nil) ||
			(got != nil && test.expectedVariable != nil && *got != *test.expectedVariable) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedVariable, got)
		}

		gotErr := err != nil
		if test.expectsErr != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %v, Line: '%s'", i, test.expectsErr, err, test.line)
		}
	}
}
func newVariable(name string, index int, initialValue uint) *data.Variable {
	variable, _ := data.NewVariable(name, index, initialValue)
	return variable
}
