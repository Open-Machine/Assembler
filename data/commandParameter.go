package data

type CommandParameter struct {
	Num   int
	Str   string
	IsStr bool
}

func NewStringParam(str string) CommandParameter {
	return CommandParameter{Num: 0, Str: str, IsStr: true}
}

func NewIntParam(num int) CommandParameter {
	return CommandParameter{Num: num, Str: "", IsStr: false}
}
