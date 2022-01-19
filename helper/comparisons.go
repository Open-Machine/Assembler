package helper

import (
	"reflect"

	"github.com/open-machine/assembler/data"
)

func SafeIsEqualStrPointer(a *string, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	return *a == *b
}

func SafeIsEqualProgramPointer(a *data.Program, b *data.Program) bool {
	return reflect.DeepEqual(a, b)
}

func SafeIsEqualInstructionPointer(a *data.Instruction, b *data.Instruction) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	return *a == *b
}
