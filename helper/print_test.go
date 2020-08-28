package helper

import (
	"github.com/open-machine/assembler/config/myerrors"
	"bytes"
	"errors"
	"fmt"
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
}

func TestPrintOut(t *testing.T) {
	stdout := Out
	defer func() { Out = stdout }()

	stderr := Err
	defer func() { Err = stderr }()

	Out = new(bytes.Buffer)
	Err = new(bytes.Buffer)
	LogInfo("a")
	stdoutStrInfo := Out.(*bytes.Buffer).String()
	stderrStrInfo := Err.(*bytes.Buffer).String()
	if !(stdoutStrInfo != "" && len(stderrStrInfo) == 0) {
		t.Errorf(fmt.Sprintf("Wrong Out in INFO. Expected STDOUT not empty and STDERR empty. STDOUT: '%s', STDERR: '%s'", stdoutStrInfo, stderrStrInfo))
	}

	Out = new(bytes.Buffer)
	Err = new(bytes.Buffer)
	PrintlnExplanation("a")
	stdoutStrExpl := Out.(*bytes.Buffer).String()
	stderrStrExpl := Err.(*bytes.Buffer).String()
	if !(stdoutStrExpl != "" && stderrStrExpl == "") {
		t.Errorf(fmt.Sprintf("Wrong Out in EXPLANATION. Expected STDOUT not empty and STDERR empty. STDOUT: '%s', STDERR: '%s'", stdoutStrExpl, stderrStrExpl))
	}
}

func TestPrintErrs(t *testing.T) {
	stdout := Out
	defer func() { Out = stdout }()

	stderr := Err
	defer func() { Err = stderr }()

	Out = new(bytes.Buffer)
	Err = new(bytes.Buffer)
	LogOtherError("a")
	stdoutStrOther := Out.(*bytes.Buffer).String()
	stderrStrOther := Err.(*bytes.Buffer).String()
	if !(stdoutStrOther == "" && stderrStrOther != "") {
		t.Errorf(fmt.Sprintf("Wrong Err in OTHER. Expected STDOUT empty and STDERR not empty. STDOUT: '%s', STDERR: '%s'", stdoutStrOther, stderrStrOther))
	}

	Out = new(bytes.Buffer)
	Err = new(bytes.Buffer)
	LogWarning("a")
	stdoutStrWarn := Out.(*bytes.Buffer).String()
	stderrStrWarn := Err.(*bytes.Buffer).String()
	if !(stdoutStrWarn == "" && stderrStrWarn != "") {
		t.Errorf(fmt.Sprintf("Wrong Err in WARNING. Expected STDOUT empty and STDERR not empty. STDOUT: '%s', STDERR: '%s'", stdoutStrWarn, stderrStrWarn))
	}

	Out = new(bytes.Buffer)
	Err = new(bytes.Buffer)
	LogErrorInLine(*myerrors.NewCodeError(errors.New("Error")), 0, "{line}")
	stdoutStrLine := Out.(*bytes.Buffer).String()
	stderrStrLine := Err.(*bytes.Buffer).String()
	if !(stdoutStrLine == "" && stderrStrLine != "") {
		t.Errorf(fmt.Sprintf("Wrong Err in LINE. Expected STDOUT empty and STDERR not empty. STDOUT: '%s', STDERR: '%s'", stdoutStrLine, stderrStrLine))
	}
}
