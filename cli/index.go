package cli

// TODO: tests

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func ConfigureCli() {
	kingpin.Version("1.0.0")

	app := kingpin.New("assembler", "Assembler command line tool.")

	ConfigureAssembleCommand(app)
	ConfigureSyntaxCommand(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
