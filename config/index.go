package config

import (
	"io"
	"os"
)

const AmntBitsCode = 8
const AmntBitsParam = 8

const VariableNameRegex = "^[a-z][a-zA-Z0-9]*$"

const AssemblyFileExtension = ".asm"

const AssemblyExampleFile = "assembly.asm"

const (
	SuccessStatus = 0
	FailStatus    = 1
)

var Out io.Writer = os.Stdout
var Err io.Writer = os.Stderr
