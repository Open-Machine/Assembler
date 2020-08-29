package helper

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/open-machine/assembler/config"
	"github.com/open-machine/assembler/config/myerrors"
)

func TestPrints(t *testing.T) {
	if config.Out != os.Stdout {
		t.Errorf("Out should be Stdout")
	}
	if config.Err != os.Stderr {
		t.Errorf("Out should be Stderr")
	}
}

func TestPrintOut(t *testing.T) {
	config.Out = new(bytes.Buffer)
	config.Err = new(bytes.Buffer)
	LogInfo("a")
	stdoutStrInfo := config.Out.(*bytes.Buffer).String()
	stderrStrInfo := config.Err.(*bytes.Buffer).String()
	if !(stdoutStrInfo != "" && len(stderrStrInfo) == 0) {
		t.Errorf(fmt.Sprintf("Wrong Out in INFO. Expected STDOUT not empty and STDERR empty. STDOUT: '%s', STDERR: '%s'", stdoutStrInfo, stderrStrInfo))
	}

	config.Out = new(bytes.Buffer)
	config.Err = new(bytes.Buffer)
	PrintlnExplanation("a")
	stdoutStrExpl := config.Out.(*bytes.Buffer).String()
	stderrStrExpl := config.Err.(*bytes.Buffer).String()
	if !(stdoutStrExpl != "" && stderrStrExpl == "") {
		t.Errorf(fmt.Sprintf("Wrong Out in EXPLANATION. Expected STDOUT not empty and STDERR empty. STDOUT: '%s', STDERR: '%s'", stdoutStrExpl, stderrStrExpl))
	}
}

func TestPrintErrs(t *testing.T) {
	config.Out = new(bytes.Buffer)
	config.Err = new(bytes.Buffer)
	LogOtherError("a")
	stdoutStrOther := config.Out.(*bytes.Buffer).String()
	stderrStrOther := config.Err.(*bytes.Buffer).String()
	if !(stdoutStrOther == "" && stderrStrOther != "") {
		t.Errorf(fmt.Sprintf("Wrong Err in OTHER. Expected STDOUT empty and STDERR not empty. STDOUT: '%s', STDERR: '%s'", stdoutStrOther, stderrStrOther))
	}

	config.Out = new(bytes.Buffer)
	config.Err = new(bytes.Buffer)
	LogWarning("a")
	stdoutStrWarn := config.Out.(*bytes.Buffer).String()
	stderrStrWarn := config.Err.(*bytes.Buffer).String()
	if !(stdoutStrWarn == "" && stderrStrWarn != "") {
		t.Errorf(fmt.Sprintf("Wrong Err in WARNING. Expected STDOUT empty and STDERR not empty. STDOUT: '%s', STDERR: '%s'", stdoutStrWarn, stderrStrWarn))
	}

	config.Out = new(bytes.Buffer)
	config.Err = new(bytes.Buffer)
	LogErrorInLine(*myerrors.NewCodeError(errors.New("Error")), 0, "{line}")
	stdoutStrLine := config.Out.(*bytes.Buffer).String()
	stderrStrLine := config.Err.(*bytes.Buffer).String()
	if !(stdoutStrLine == "" && stderrStrLine != "") {
		t.Errorf(fmt.Sprintf("Wrong Err in LINE. Expected STDOUT empty and STDERR not empty. STDOUT: '%s', STDERR: '%s'", stdoutStrLine, stderrStrLine))
	}
}
