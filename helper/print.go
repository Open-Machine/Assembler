package helper

import (
	"fmt"
	"io"
	"os"
)

var Out io.Writer = os.Stdout
var Err io.Writer = os.Stderr

func PrintOut(str string) {
	fmt.Fprint(Out, str)
}

func PrintlnOut(str string) {
	fmt.Fprint(Out, str+"\n")
}

func PrintErr(str string) {
	fmt.Fprint(Err, str)
}

func PrintlnErr(str string) {
	fmt.Fprint(Err, str+"\n")
}
