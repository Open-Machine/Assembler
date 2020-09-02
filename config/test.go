package config

import (
	"math"
	"testing"
)

var Testing bool = false

func TestSetup() {
	Testing = true
}

func BitsRelation(t *testing.T) {
	expectedAmntHexDigitsInstruct := math.Log2(AmntBitsInstruction)
	if expectedAmntHexDigitsInstruct != AmntHexDigitsInstruction {
		t.Errorf("Expected %f hex digits in instruction, but got %d", expectedAmntHexDigitsInstruct, AmntHexDigitsInstruction)
	}

	expectedAmntHexDigitsParam := math.Log2(AmntBitsParam)
	if expectedAmntHexDigitsParam != AmntBitsParam {
		t.Errorf("Expected %f hex digits in param, but got %d", expectedAmntHexDigitsParam, AmntHexDigitsParam)
	}
}
