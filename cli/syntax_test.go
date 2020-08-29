package cli

import (
	"strings"
	"testing"
)

// TODO: test if the shell instructions are returning the right AssembleInstruction
// TODO: test getSyntaxExample() -> check if the program compiles

func TestAssemblyCompiledExample(t *testing.T) {
	fileString, err := getAssemblyExample()

	if err != nil {
		t.Errorf("Could not open file. Err: '%s'", err.Error())
	}

	replaced := strings.ReplaceAll(fileString, " ", "")
	replaced = strings.ReplaceAll(fileString, "\n", "")
	replaced = strings.ReplaceAll(fileString, "\r", "")
	if replaced == "" {
		t.Errorf("Blank file")
	}
}
