<h1 align="center">Open Machine's Assembler</h1>
<div align="center"> 

<a href="https://github.com/Open-Machine/Assembler/stargazers"><img src="https://img.shields.io/github/stars/Open-Machine/Assembler" alt="Stars Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/network/members"><img src="https://img.shields.io/github/forks/Open-Machine/Assembler" alt="Forks Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/"><img src="https://img.shields.io/badge/version-0.0.1-blue" alt="Version Badge"/></a>
<a href="https://coveralls.io/github/Open-Machine/Assembler"><img src="https://img.shields.io/coveralls/github/Open-Machine/Assembler" alt="coverage"></a>
<a href="https://github.com/Open-Machine/Assembler/commits/"><img src="https://img.shields.io/github/commit-activity/m/Open-Machine/Assembler" alt="commits"/></a>
<a href="https://github.com/Open-Machine/Assembler/pulls"><img src="https://img.shields.io/github/issues-pr/Open-Machine/Assembler" alt="Pull Requests Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/issues"><img src="https://img.shields.io/github/issues/Open-Machine/Assembler" alt="Issues Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Open-Machine/Assembler?color=2b9348"></a>
<a href="https://github.com/Open-Machine/Circuits/blob/master/LICENSE"><img src="https://img.shields.io/github/license/Open-Machine/Assembler?color=2b9347" alt="License Badge"/></a>

<h3><i><b>Assembler</b> of Open-Machine's <a href="https://github.com/Open-Machine/Circuits">Circuit</a>. The program reads assembly code and generates machine code.</i></h3>
<br/>

<h3><i>This repository is part of a <a href="https://github.com/Open-Machine/">larger project</a>: <b>developing a computer from scratch</b> with its custom <a href="https://github.com/Open-Machine/Circuits">circuit</a>, <a href="https://github.com/Open-Machine/Assembler">assembler</a> and compiler.</i></h3>
<br/>

<img src="https://raw.githubusercontent.com/Open-Machine/README/stable/Media/logo-horizontal.png" alt="open-machine"/>

<br/>

</div>

---

# 🔖 Table of Contents

## Introduction
### 1. [📌 Definition and Explanation](#-definition-and-explanation)

## Code
### 2. [🔢 Instructions](#-instructions)
### 3. [🔀 Code Flow and Tips](#-code-flow-and-tips)
### 4. [🔡 Code syntax](#-code-syntax)
### 5. [👨🏻‍💻 Code Example](#-code-example)

## Run
### 6. [:arrow_forward: Setup and Run](#arrow_forward-setup-and-run)
### 7. [💻 Assembler CLI](#-assembler-cli)

## More
### 9. [📄 Contributing Guidelines](#-contributing-guidelines)

---

# 📌 Definition and Explanation
Assembly is basically the most basic programming language of a certain hardware. There's a very strong correspondence between the instructions in the language and the architecture's machine code instructions: every instruction in the language is a machine code instruction and vice-versa. It was created so that humans don't had to memorize the machine code instructions which are many numbers.

From the Wikipedia:
> In computer programming, assembly language, often abbreviated asm, is any low-level programming language in which there is a very strong correspondence between the instructions in the language and the architecture's machine code instructions.

Because of this strong correspondence, the translating process is called assembling instead of compiling, which is the same process but for high-end languages. Those languages do not have this strong correspondence that assembly languages have.

The core of the assembling process is to identify the assembly instructions and translate them to the circuit's instruction binary equivalent. Similarly, it also has to convert each variable to a memory address.

## Learn more

If you are interested in knowing more how this process works don't be afraid to read at the code.

If you are interested in knowing more about the actual circuit that runs the code you write, click [here](https://github.com/Open-Machine/Circuits/).

If you are interested in knowing more about the Open-Computer project, click [here](https://github.com/Open-Machine/README/).

---

# 🔢 Instructions
Let's take a close look at the built-in instructions available. **Don't worry about syntax right now, we will talk about it later.**

**Warning** ⚠️: If you have never programmed in an assembly language, please read the section [🔀 Code Flow and Tips](#-code-flow-and-tips).

### Symbols Legend
Some symbols are used in the Instructions Table. Here you can see their meaning.
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

---

# 🔀 Code Flow and Tips

Because Open-Machine's Circuit only has very simple commands and few registers, the way to think about your assembly code will be very different.

### Storage
The circuit has two components that store data: the ACC register and the memory RAM. Both of these are volatile memories, which means that when the circuit is turned the data lost. Let's take a closer look in the memories available and when they should be used:
- **RAM**: is the main memory, it can store thousands of bits and every variable should be stored here.
- **ACC register**: is an auxiliary memory for arithmetic operations that can store only one value. It must not store variables indefinitely. 

	Most CPUs have many registers, so in those cases some registers can be used to store variables indefinitely. However, since Open-Computer's circuit only offers one register, it must be used exclusively as an auxiliary memory for the instructions.

### Operation Flow
Since the circuit has only one register, the flow of the operations will be a little bit different, following a pattern somewhat similar to:

1. Change the value of ACC register
2. Do an instruction
3. Store the value of the ACC register in RAM

**For example**, if you want to sum variables A and B and store the result in C you could use the following instructions:
1. Copy the value of variable A to the ACC register
2. Use the sum instruction to sum the value of ACC register with B and store the result in the ACC register
3. Store the value of the ACC register in C memory address

### IFs, WHILEs, FORs and procedures
If that's your first time programming assembly, it must be very strange to know that there are no ```if```s, ```while```s and ```for```s. However, it's not that hard not having those keywords, because all of those things can be done with the combination arithmetic operations and conditional and unconditional jumps.

Let me show you an example. Imagine you have this code written in C and wanted to translate it to assembly.
```c
	// before
	if (a > b) {
		// ...
	}
	// after
```
One way of doing it would be:
1. **Copy** the value of ```a``` from RAM to the ACC register
2. Update the value of the ACC register with the result of the ```subtraction``` between the value of the ACC register and ```b```
3. **Jump** to **step 5** if the ACC register is greater than zero
4. One or more instructions inside the ```if``` statement
	
	```c
	// after
	```

5. After instructions

	```c
	// after
	```

## More Tips
- Remember to add the ```kill``` instruction at the end of your programs

---

# 🔡 Code Syntax
**Warning** ⚠️: Assembly languages are specific to their hardware so remember that Open-Computer's Assembly may be different from other assembly languages.

Read the specifications below to learn the code syntax.

## Tabs, spaces and case sensitivity
- Case sensitive;
- **Tabs and spaces** can be used interchangeably;
- Blank or empty lines won't be considered;
- **Numbers** can be written in hexadecimal in the form of ```0xff``` or in decimal as ```255```;

## Naming Practices
- A **label name** should start with a letter and the rest of the name can have more letters and numbers;
- Every name should obey the following regex: ```[a-z][a-zA-Z0-9]*```;
- Snake-case is not allowed and the use of camel-case is encouraged.

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
Check out [here](#-instructions) the instruction table to know what instructions you can use and their parameters.

---

# 👨🏻‍💻 Code Example
```sh
# Let's get two numbers as input and print the result of the sum of them



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

# :arrow_forward: Setup and Run
These are the steps to setup and run Open-Computer's Assembler.

You can find more information about the assembler CLI [here](#-assembler-cli) and about running the circuit [here](https://github.com/Open-Machine/Circuits/#️-run).

## Setup
1. Build the GoLang project
	```sh
	go build
	```
2. Clone [Open-Computer's Circuit Repository](https://github.com/Open-Machine/Assembler/)

	You will need this repository to run the assembled program.

	If you have git installed in your terminal, run:
	```sh
	git clone https://github.com/Open-Machine/Assembler/
	```

## Assemble your code
1. **Assemble** your code
   ```sh
   ./assembler assemble ${main.asm}
   ```

## Run your code   
There are two ways of running your application

#### GUI Mode
In this mode, you will be able to see everything that is happening to the circuits in real time and interact with it by changing the inputs.

You can watch [this video](https://www.youtube.com/watch?v=NAITQqdOw7c) as an introduction to Logisim-Evolution, which is the program that we will be using to simulate the circuit.

1. Navigate to the Circuits repository
2. Start the circuit: follow the steps [here](https://github.com/Open-Machine/Circuits/#️start-the-circuit)
3. Right click in the RAM and click ""
4. Select the assembled file

   You should select the file generated by the assemble program, not the file with the assembly code.

5. To run the program (start the clock simulation), follow the steps [here](https://github.com/Open-Machine/Circuits/#run-the-circuit)


## CLI Mode
In this mode, you will only be able to see the outputs of your application. You just have to run: 
```sh
java -jar logisim-evolution.jar main.circ -load ${assembled_file} -tty table
```
*Remember to write the name of the file that was generated by the assembler command instead of ```${assembled_file}```*.

### About the outputs

The outputs will appear on the console following this pattern: ```{16 bits of the main output}     {4 bit ouptut counter}```.

The first output can be ignored.

---

## 💻 Assembler CLI
You can use the flag ```--help``` to see all the options.

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

# 📄 Contributing Guidelines
Check out the contributing guidelines [here](https://github.com/Open-Machine/github.com/open-machine/assemblerblob/master/CONTRIBUTION.md).
