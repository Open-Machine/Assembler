package config

import (
	"io"
	"os"
)

const AmntBitsInstruction = 4
const AmntHexDigitsInstruction = 1
const AmntBitsParam = 12
const AmntHexDigitsParam = 3

const VariableNameRegex = "^[a-z][a-zA-Z0-9]*$"

const AssemblyFileExtension = ".asm"
const MachineCodeFileExtension = ".run"

const RunMachineCodeScriptPath = "run_machine_code.sh"

const (
	SuccessStatus = 0
	FailStatus    = 1
)

var Out io.Writer = os.Stdout
var Err io.Writer = os.Stderr
