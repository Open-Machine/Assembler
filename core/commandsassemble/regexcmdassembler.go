package commandsassembler

type RegexCommandAssembler struct {
	regex    string
	Assemble funcAssemble
}

func NewRegexCommandAssembler(regex string, assemble funcAssemble) RegexCommandAssembler {
	return RegexCommandAssembler{regex: regex, Assemble: assemble}
}
