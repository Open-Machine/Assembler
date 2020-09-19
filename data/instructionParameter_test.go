package data

import (
	"testing"

	"github.com/open-machine/assembler/config"
)

func TestStrParamConstructors(t *testing.T) {
	var tests = []struct {
		str         string
		expectedErr bool
	}{
		// Success
		{
			str:         "helloWorld",
			expectedErr: false,
		},
		// Fail
		{
			str:         "Hello World",
			expectedErr: true,
		},
		{
			str:         "Hello",
			expectedErr: true,
		},
	}

	for i, test := range tests {
		strParam, err := NewStringParam("Hello World")
		gotErr := err != nil

		if gotErr != test.expectedErr {
			t.Errorf("[%d] Expected error: %t, got error: %t", i, test.expectedErr, gotErr)
		}

		if !strParam.IsStr {
			t.Errorf("String param should have true isStr")
		}
	}
}

func TestIntParamConstructors(t *testing.T) {
	var tests = []struct {
		num         int
		expectedErr bool
	}{
		// Success
		{
			num:         1,
			expectedErr: false,
		},
		// Fail
		{
			num:         5000,
			expectedErr: true,
		},
	}

	for i, test := range tests {
		strParam, err := NewStringParam("Hello World")
		gotErr := err != nil

		if gotErr != test.expectedErr {
			t.Errorf("[%d] Expected error: %t, got error: %t", i, test.expectedErr, gotErr)
		}

		if strParam.IsStr {
			t.Errorf("Int param should have false isStr")
		}
	}
}

func TestParamConstructorsTesting(t *testing.T) {
	num := 0
	str := "a"
	isStr := true

	config.Testing = true
	instruc1 := NewParamTest(num, str, isStr)
	if instruc1 == nil {
		t.Errorf("Instruction should not be nil")
	}

	config.Testing = false
	instruc2 := NewParamTest(num, str, isStr)
	if instruc2 != nil {
		t.Errorf("Instruction should be nil")
	}
}
