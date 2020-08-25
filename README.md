# Computer-from-Scratch
Designing a computer from scratch. The only components that I didn't made were the ports.

#### Todo
- [ ] High-Level Language Compiler
- [ ] Signed integer
- [ ] Float
- [ ] Division circuit and command
- [ ] Clock
- [ ] Own RAM

## Objective
Understand how a computer works behind the curtains and maybe do some things my way.

---

# 🔖 Table of Contents
- [👨🏻‍💻 Coding](#coding)
- [▶️ Assembling (shell commands)](#assembling)

---

## Coding
From the Wikipedia:
> In computer programming, assembly language, often abbreviated asm, is any low-level programming language in which there is a very strong correspondence between the instructions in the language and the architecture's machine code instructions.

Assembly is basically the most basic programming language of a certain hardware. There's a very strong correspondence between the instructions in the language and the architecture's machine code instructions: every instruction in the language is a machine code instruction and vice-versa.

In this case the translating process is called assembling, because ...

The code should be written in RAM and it will be executed from the memmory address in register CP content. Every 4 bytes are considered a line of code.

Line of code = Instruction (2 bytes) + Memory Address (2 bytes).

ps: The memory address in the lines of code will be called EE - [EE] represents EE value and EE

###### Instructions Table
Machine Code | Assembly Command | Description | Requires Param
--- | --- | --- | ---
0x0 | nop | No operation | No
0x1 | copy | [AC] = [EE] | Yes
0x2 | store | [EE] = [AC] | Yes
0x3 | add | [AC] = [AC] + [EC] | Yes
0x4 | sub | [AC] = [AC] - [EC] | Yes
0x7 | input | [EE] = to the input value | Yes
0x8 | output | Output [EE] | Yes
0x9 | kill | Finish program | No
0xa | jmp | Jump to EE | Yes
0xb | jg | Jump to EE if [AC] > 0 | Yes
0xd | je | Jump to EE if [AC] = 0 | Yes
0xf | jl | Jump to EE if [AC] < 0 | Yes
### Machine Code Example
```sh
01ff # copy value in the address ff in RAM
020a # stores the value of AC in the address 0a
0900 # kills program
```
### Assembly Code Example
```sh
copy 0xff # copy value in the address ff in RAM
store 0x0a # stores the value of AC in the address 0a
kill
```

---

## Assembling
```sh
./assembler assemble filename.asm
```
