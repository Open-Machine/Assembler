package helper

import (
	"assembler/data"
	"reflect"
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
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	return reflect.DeepEqual(*a, *b)
}

func SafeIsEqualCommandPointer(a *data.Command, b *data.Command) bool {
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

func SafeIsEqualCommandParamPointer(a *data.CommandParameter, b *data.CommandParameter) bool {
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
