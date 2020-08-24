package helper

import (
	"bytes"
	"os"
	"testing"
)

func TestPrints(t *testing.T) {
	if Out != os.Stdout {
		t.Errorf("Out should be Stdout")
	}
	if Err != os.Stderr {
		t.Errorf("Out should be Stderr")
	}

	stdout := Out
	Out = new(bytes.Buffer)
	defer func() { Out = stdout }()

	stderr := Err
	Err = new(bytes.Buffer)
	defer func() { Err = stderr }()

	PrintOut("a")
	PrintlnOut("b")
	stdoutStr := Out.(*bytes.Buffer).String()
	if stdoutStr != "ab\n" {
		t.Errorf("Wrong Out")
	}

	PrintErr("a")
	PrintlnErr("b")
	stderrStr := Err.(*bytes.Buffer).String()
	if stderrStr != "ab\n" {
		t.Errorf("Wrong Err")
	}
}
