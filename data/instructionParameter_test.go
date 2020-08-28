package data

import "testing"

func TestParamConstructors(t *testing.T) {
	strParam := NewStringParam("Hello World")
	if !strParam.IsStr {
		t.Errorf("String param should have true isStr")
	}

	numParam := NewIntParam(1)
	if numParam.IsStr {
		t.Errorf("Int param should have false isStr")
	}
}
