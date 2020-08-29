package utils

import (
	"bytes"
	"testing"
)

func TestMyBufferAsFile(t *testing.T) {
	buffer := new(bytes.Buffer)
	name := "file.asm"
	myFile := NewMyBufferAsFile(buffer, name)

	if myFile.Reader() != buffer {
		t.Errorf("Different buffer")
	}

	gotName := myFile.Name()
	if gotName != name {
		t.Errorf("Different name. Expected: '%s', Got: '%s'", name, gotName)
	}

	if myFile.Close() != nil {
		t.Errorf("Expected nil error")
	}
}
