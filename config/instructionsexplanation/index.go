package instructionsexplanation

type InstructionExplanation struct {
	Instruction string
	Param       string
}

var instructionExplanations = map[string]InstructionExplanation{
	"nop": {
		Instruction: "No operation",
		Param:       "No param needed",
	},
	"copy": {
		Instruction: "Copies the value from memory to the ACC register",
		Param:       "The parameter refers to the memory address",
	},
	"store": {
		Instruction: "Stores the value of the ACC register in memory",
		Param:       "The parameter refers to the memory address",
	},
	"add": {
		Instruction: "Adds a memory value to the ACC register and stores the result in ACC",
		Param:       "The parameter refers to the memory address",
	},
	"sub": {
		Instruction: "Subtracts a memory value from the value of the ACC register and stores the result in ACC",
		Param:       "The parameter refers to the memory address",
	},
	"input": {
		Instruction: "Inputs the input value into memory",
		Param:       "The parameter refers to the memory address",
	},
	"output": {
		Instruction: "Outputs a memory value",
		Param:       "The parameter refers to the memory address",
	},
	"kill": {
		Instruction: "Kills the program (you will need this instruction to tell the computer that the program ended)",
		Param:       "No param needed",
	},
	"jmp": {
		Instruction: "Jumps to the a instruction",
		Param:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
	"jg": {
		Instruction: "Jumps to the a instruction if ACC register is greater than zero",
		Param:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
	"je": {
		Instruction: "Jumps to the a instruction if ACC register is equal to zero",
		Param:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
	"jl": {
		Instruction: "Jumps to the a instruction if ACC register is less than zero",
		Param:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
}
