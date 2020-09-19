<h1 align="center">Open Machine's Assembler</h1>
<div align="center">

<a href="https://github.com/Open-Machine/Assembler/stargazers"><img src="https://img.shields.io/github/stars/Open-Machine/Assembler" alt="Stars Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/network/members"><img src="https://img.shields.io/github/forks/Open-Machine/Assembler" alt="Forks Badge"/></a>
<a href="https://coveralls.io/github/Open-Machine/Assembler"><img src="https://img.shields.io/coveralls/github/Open-Machine/Assembler" alt="coverage"></a>
<a href="https://github.com/Open-Machine/Assembler/commits/"><img src="https://img.shields.io/github/commit-activity/m/Open-Machine/Assembler" alt="commits"/></a>
<a href="https://github.com/Open-Machine/Assembler/pulls"><img src="https://img.shields.io/github/issues-pr/Open-Machine/Assembler" alt="Pull Requests Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/issues"><img src="https://img.shields.io/github/issues/Open-Machine/Assembler" alt="Issues Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Open-Machine/Assembler?color=2b9348"></a>
<a href="https://github.com/Open-Machine/Assembler/blob/master/LICENSE"><img src="https://img.shields.io/github/license/Open-Machine/Assembler?color=2b9347" alt="License Badge"/></a>

<img src="https://raw.githubusercontent.com/Open-Machine/README/master/Media/logo-horizontal.png" alt="open-machine"/>

<br/>
<i><b>Assembler</b> of Open-Machine's <a href="https://github.com/Open-Machine/Circuits">Circuit</a>. The program reads assembly code and generates machine code.</i>
<br/>

<i>This repository is part of a <a href="https://github.com/Open-Machine/">larger project</a>: <b>developing a computer from scratch</b> with its custom <a href="https://github.com/Open-Machine/Circuits">circuit</a>, <a href="https://github.com/Open-Machine/Assembler">assembler</a> and compiler.</i>

</div>

---

# 🔖 Table of Contents
### 1. [✔ Todo](#-todo)
### 2. [📌 Definition and Explanation (WIP)](#-definition-and-explanation-wip)
### 3. [🔡 Code syntax](#-code-syntax)
### 4. [👨🏻‍💻 Code Example](#-code-example)
### 5. [:arrow_forward: Assembler CLI](#arrow_forward-assembler-cli)
### 6. [:bug: Build and Test](#bug-build-and-test)
### 7. [📄 Contributing Guidelines](#-contributing-guidelines)

---

# ✔ Todo
- [X] Core
- [X] Add jump labels
- [X] Add comments
- [ ] Variables
- [ ] Import
- [ ] Procedures

---

# 📌 Definition and Explanation (WIP)
Assembly is basically the most basic programming language of a certain hardware. There's a very strong correspondence between the instructions in the language and the architecture's machine code instructions: every instruction in the language is a machine code instruction and vice-versa. It was created so that humans don't had to memorize the machine code instructions which are many numbers.

From the Wikipedia:
> In computer programming, assembly language, often abbreviated asm, is any low-level programming language in which there is a very strong correspondence between the instructions in the language and the architecture's machine code instructions.

In this case the translating process is called assembling, because ...

The code should be written in RAM and it will be executed from the memmory address in register CP content. Every 4 bytes are considered a line of code.

## Memory

Line of code = Instruction (4 bits) + Memory Address (12 bits).

For more information about the machine code and the circuit, check out the [Circuits Repository](https://github.com/Open-Machine/Circuits/).

---

# 🔡 Code Syntax
This is the assembly of Open-Machine's Circuits and it may be different from assemblies of other computers.

The circuit has two types of memories that can be used to store data: variables, which will be store in the RAM, and the ACC register, which should be used as an auxiliary memory for arithmetic operations.

Read the specifications below to learn how to code in Open-Machine's Assembly.

## Tabs, spaces and case sensitivity
- Case sensitive;
- **Tabs and spaces** can be used interchangeably;
- Blank or empty lines will be disconsidered;
- **Numbers** can be written in hexadecimal in the form of ```0xff``` or in decimal as ```255```;

## Naming Practices
- A **variable, procedure or label name** should start with a letter and the rest of the name can have more letters and numbers;
- Every name should obey the following regex: ```[a-z][a-zA-Z0-9]*```;
- Snake-case is not allowed and the use of camel-case is encouraged.

## Imports
- The important files cannot have any code outside procedures.
- Form: ```import {fileName}.asm```
- Example:
	```sh
	import help.asm
	```

## Declaring Variables
- The variables should be declared between the imports and instructions.
- Form: ```declare {variableName} {number}```
- Example
	```sh
	declare variable 0xff
	```
- Remember to follow the [naming practices](#naming-practices)

## Jump Label
- Definition: it marks the line for possible jumps to that line;
- Form: ```{labelName}:```
- Remember to follow the [naming practices](#naming-practices)

## Instruction line
***Definition***

An instruction line is a line that contains an instruction call.

***Components***
- ```instruction``` is the actual instruction that will be executed, it must be one of the following in the instruction table;
- ```arg``` can be either a jump label or a 

***Form***
- A instruction line should be in the following form ```{instruction} [arg]```;
- An instruction line should obey the following regex: ```^[\t ]*(((nop)|(copy)|(store)|(add)|(sub)|(input)|(output)|(kill)|(jmp)|(jg)|(je)|(jl))(([\t ]+[a-z][a-zA-Z0-9]*)|()))[\t ]*$```

## Instructions Table
### Symbols Legend for the Instructions Table
Symbol | Explanation
--- | ---
ACC | The ACC register
variable | A variable from the memory
label | Jump label
[ ] | "Value of"
### Instructions Table
Assembly Command | Short Instruction Description | Long Instruction Description | Short Param Description | Long Param Description
--- | --- | --- | --- | ---
nop | - | This instruction doesn't perform any action | - | No parameter is required
copy | [ACC] = [variable] | A value from the memory is copied to the ACC register | variable | It's the name of the variable that will be used in the instruction
store | [variable] = [ACC] | The value from the ACC register is stored into memory | variable | It's the name of the variable that will be used in the instruction
add | [ACC] = [ACC] + [variable] | The sum of the value of the ACC register and a value from the memory is stored in the ACC register | variable | It's the name of the variable that will be used in the instruction
sub | [ACC] = [ACC] - [variable] | The difference between the value of the ACC register and a value from the memory is stored in the ACC register | variable | It's the name of the variable that will be used in the instruction
input | [variable] = input value | The input value is copied to the memory | variable | It's the name of the variable that will be used in the instruction
output | Output [variable] | Outputs a value from the memory into the circuit LEDs | variable | It's the name of the variable that will be used in the instruction
kill | Finishes program | When this instruction is encountered, the program is finished and no more instructions will be executed | - | No parameter is required
jmp | Jump to EE | Jump to another line of code | label | The jump label the program will jump to
jg | Jump to EE if [ACC] > 0 | Jump to another line of code if the value of the ACC register is positive | label | The jump label the program will jump to if the condition is right
je | Jump to EE if [ACC] = 0 | Jump to another line of code if the value of the ACC register is zero | label | The jump label the program will jump to if the condition is right
jl | Jump to EE if [ACC] < 0 | Jump to another line of code if the value of the ACC register is negative | label | The jump label the program will jump to if the condition is right

## Procedures
```sh
procedure procSum
	store variable
	copy variable
end
```

## Tips
- Remember to add the ```kill``` instruction at the end of your programs

---

# 👨🏻‍💻 Code Example
```sh
import help.asm # it has procedure procSum

procedure procSub
	copy 0xff # copy value in the address ff in RAM
	store 0x0a # stores the value of ACC in the address 0a
end

start:
	procSub

	copy 0xff # copy value in the address ff in RAM
	store 0x0a # stores the value of ACC in the address 0a
	je start # jumps to the 'start' label if ACC is 0
kill
```

---

# :arrow_forward: Assembler CLI
You can use the flag ```--help``` to see the options 

### Assemble
```sh
./assembler assemble --help
```
```
usage: assembler assemble [<flags>] <file-name>

Assemble assembly code

Flags:
      --help            Show context-sensitive help (also try --help-long and
                        --help-man).
  -r, --rename-exec=""  Provide the name of the executable file that will be created
                        (if empty, the name will be the same as the assembly code
                        file)

Args:
  <file-name>  Provide the name of file with the assembly code
```
*ps: If you want to **assemble your assembly file** run:* ```./assembler assemble filename.asm```

### Syntax
```sh
./assembler syntax --help
```
```
usage: assembler syntax [<flags>]

Help with the syntax of this assembly language

Flags:
      --help                     Show context-sensitive help (also try --help-long
                                 and --help-man).
  -e, --example                  Assembly code example with explanation
  -l, --ls                       List all available instructions
  -c, --instruction=INSTRUCTION  Explanation of an specific instruction
```

---

# :bug: Build and Test

## Build
To build the GoLang code run:
```sh
go build
```

## Test
To test the GoLang code run:
```sh
go test ./...
```

---

# 📄 Contributing Guidelines
Check out the contributing guidelines [here](https://github.com/Open-Machine/github.com/open-machine/assemblerblob/master/CONTRIBUTION.md).
