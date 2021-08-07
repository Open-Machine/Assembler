package cli

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func ConfigureCli() {
	kingpin.Version("0.0.2")

	app := kingpin.New("assembly", "Assembler instruction line tool.")

	ConfigureAssembleCommand(app)
	ConfigureSyntaxCommand(app)
	ConfigureRunCommand(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
