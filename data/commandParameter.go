package data

type InstructionParameter struct {
	Num   int
	Str   string
	IsStr bool
}

func NewStringParam(str string) InstructionParameter {
	return InstructionParameter{Num: 0, Str: str, IsStr: true}
}

func NewIntParam(num int) InstructionParameter {
	return InstructionParameter{Num: num, Str: "", IsStr: false}
}
