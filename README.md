<h1 align="center">Assembler - from Open Machine</h1>
<div align="center">

<a href="https://github.com/Open-Machine/Assembler/stargazers"><img src="https://img.shields.io/github/stars/Open-Machine/Assembler" alt="Stars Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/network/members"><img src="https://img.shields.io/github/forks/Open-Machine/Assembler" alt="Forks Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/pulls"><img src="https://img.shields.io/github/issues-pr/Open-Machine/Assembler" alt="Pull Requests Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/issues"><img src="https://img.shields.io/github/issues/Open-Machine/Assembler" alt="Issues Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Open-Machine/Assembler?color=2b9348"></a>
<a href="https://github.com/Open-Machine/Circuits/blob/master/LICENSE"><img src="https://img.shields.io/github/license/Open-Machine/Assembler?color=2b9348" alt="License Badge"/></a>

<img src="https://raw.githubusercontent.com/Open-Machine/README/master/Media/logo-horizontal.png" alt="open-machine"/>

<br/>
<i>The goal is to to design and build a computer from scratch only using logical gates to do so.<i/>
<br/>

<i>This repository is part of a <a href="https://github.com/Open-Machine/">larger project<a/>: developing a computer from scratch with its <a href="https://github.com/Open-Machine/Assembler">assembler<a/> and compiler.

</div>

# Assembler
A program that transforms assembly code into machine code.

*This repository is part of a bigger project: **developing an open-source computer from scratch**. Check out **[Open Machine](https://github.com/Open-Machine/)**!*

#### Todo
- [ ] Variables
- [ ] Import
- [ ] Procedures

## Objective
Understand how a computer works behind the curtains and maybe do some things my way.

---

# 🔖 Table of Contents
### 1. [📌 Purpose and Definition](#purpose-and-definition)
### 2. [🔡 Code syntax](#code-syntax)
### 3. [👨🏻‍💻 Code Example](#example)
### 3. [▶️ Assembler CLI](#assembler-cli)
### 4. [📄 Contributing Guidelines](#contributing-guidelines)

---

# 📌 Purpose and Definition
Assembly is basically the most basic programming language of a certain hardware. There's a very strong correspondence between the instructions in the language and the architecture's machine code instructions: every instruction in the language is a machine code instruction and vice-versa. It was created so that humans don't had to memorize the machine code instructions which are many numbers.

From the Wikipedia:
> In computer programming, assembly language, often abbreviated asm, is any low-level programming language in which there is a very strong correspondence between the instructions in the language and the architecture's machine code instructions.

In this case the translating process is called assembling, because ...

The code should be written in RAM and it will be executed from the memmory address in register CP content. Every 4 bytes are considered a line of code.

## Memory

Line of code = Instruction (2 bytes) + Memory Address (2 bytes).

ps: The memory address in the lines of code will be called EE - [EE] represents EE value and EE

For more information about the machine code and the circuit, check out the [Circuits Repository]().

---

# 🔡 Code Syntax
This assembly 

RAM Memory and ACC

Sum, subtract, store and , jump to another instruction

## Some Specifications

#### Tabs, spaces and cases 
- Case sensitive;
- **Tabs and spaces** can be used interchangeably;
- Blank or empty lines will be disconsidered;
- **Numbers** can be written in hexadecimal in the form of ```0xff``` or in decimal as ```255```;

#### Naming Practices
A **variable, procedure or label name** should start with a letter and the rest of the name can have more letters and numbers;
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
input | [variable] = to the input value | The input value is copied to the memory | variable | It's the name of the variable that will be used in the instruction
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

# ▶️ Assembler CLI
You can use the flag ```--help``` to see the options 

#### Assemble code
**Assemble file**
```sh
./assembler assemble filename.asm
```
**Help**
```sh
./assembler assemble --help
```
```

```

#### Syntax explanation
```sh
./assembler syntax --help
```
```

```

---

# 📄 Contributing Guidelines
Check out the contributing guidelines [here](https://github.com/Open-Machine/github.com/open-machine/assemblerblob/master/CONTRIBUTION.md).
