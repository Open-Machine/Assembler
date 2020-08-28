package instructionsexplanation

type InstructionExplanation struct {
	Instruction string
	Param       string
}

var instructionExplanations = map[string]InstructionExplanation{
	"nop": InstructionExplanation{
		Instruction: "No operation",
		Param:       "No param needed",
	},
	"copy": InstructionExplanation{
		Instruction: "Copies the value from memory to the ACC register",
		Param:       "The parameter refers to the memory address",
	},
	"store": InstructionExplanation{
		Instruction: "Stores the value of the ACC register in memory",
		Param:       "The parameter refers to the memory address",
	},
	"add": InstructionExplanation{
		Instruction: "Adds a memory value to the ACC register and stores the result in ACC",
		Param:       "The parameter refers to the memory address",
	},
	"sub": InstructionExplanation{
		Instruction: "Subtracts a memory value from the value of the ACC register and stores the result in ACC",
		Param:       "The parameter refers to the memory address",
	},
	"input": InstructionExplanation{
		Instruction: "Inputs the input value into memory",
		Param:       "The parameter refers to the memory address",
	},
	"output": InstructionExplanation{
		Instruction: "Outputs a memory value",
		Param:       "The parameter refers to the memory address",
	},
	"kill": InstructionExplanation{
		Instruction: "Kills the program (you will need this instruction to tell the computer that the program ended)",
		Param:       "No param needed",
	},
	"jmp": InstructionExplanation{
		Instruction: "Jumps to the a instruction",
		Param:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
	"jg": InstructionExplanation{
		Instruction: "Jumps to the a instruction if ACC register is greater than zero",
		Param:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
	"je": InstructionExplanation{
		Instruction: "Jumps to the a instruction if ACC register is equal to zero",
		Param:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
	"jl": InstructionExplanation{
		Instruction: "Jumps to the a instruction if ACC register is less than zero",
		Param:       "The parameter can be either a label or a number that refers to a instruction (Warning: the index of a instruction can be different from the index of the line where the instruction is)",
	},
}
