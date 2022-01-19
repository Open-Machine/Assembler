<div align="center"> 
<h1>Open Machine's Assembler</h1>
<h4>CLI app that reads Assembly code and generates Machine Code for Open-Machine's Circuit.</h4>

<i>This repository is a component of a larger project: <b><a href="https://github.com/Open-Machine/README">Open-Machine</a></b> - an open-source computer developed from scratch.</i>

<b>⚠️ It relies heavily on the <a href="https://github.com/Open-Machine/Circuits">Circuits</a> repository and it will not run unless both are in the same folder with the repository names unchanged.</b>

<a href="https://github.com/Open-Machine/Assembler/stargazers"><img src="https://img.shields.io/github/stars/Open-Machine/Assembler" alt="Stars Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/network/members"><img src="https://img.shields.io/github/forks/Open-Machine/Assembler" alt="Forks Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/"><img src="https://img.shields.io/badge/version-0.0.1-blue" alt="Version Badge"/></a>
<a href="https://coveralls.io/github/Open-Machine/Assembler"><img src="https://img.shields.io/coveralls/github/Open-Machine/Assembler" alt="coverage"></a>
<a href="https://github.com/Open-Machine/Assembler/commits/"><img src="https://img.shields.io/github/commit-activity/m/Open-Machine/Assembler" alt="commits"/></a>
<a href="https://github.com/Open-Machine/Assembler/pulls"><img src="https://img.shields.io/github/issues-pr/Open-Machine/Assembler" alt="Pull Requests Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/issues"><img src="https://img.shields.io/github/issues/Open-Machine/Assembler" alt="Issues Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Open-Machine/Assembler?color=2b9348"></a>
<a href="https://github.com/Open-Machine/Circuits/blob/master/LICENSE"><img src="https://img.shields.io/github/license/Open-Machine/Assembler?color=2b9347" alt="License Badge"/></a>

<br/>

<img src="https://raw.githubusercontent.com/Open-Machine/README/stable/Media/logo-horizontal.png" alt="open-machine"/>

<br/>

</div>

<br/>

---

<br/>

# 🔖 Table of Contents

##### *Introduction*
### &nbsp;&nbsp;&nbsp;&nbsp;1. [📌 Definition and Explanation](#-definition-and-explanation)

##### *Code*
### &nbsp;&nbsp;&nbsp;&nbsp;2. [🔢 Instructions](#-instructions)
### &nbsp;&nbsp;&nbsp;&nbsp;3. [🔀 Code Flow and Tips](#-code-flow-and-tips)
### &nbsp;&nbsp;&nbsp;&nbsp;4. [🔡 Code syntax](#-code-syntax)
### &nbsp;&nbsp;&nbsp;&nbsp;5. [⌨️ Code Example](#-code-example)

##### *Run*
### &nbsp;&nbsp;&nbsp;&nbsp;6. [:arrow_forward: Setup and Run](#arrow_forward-setup-and-run)
### &nbsp;&nbsp;&nbsp;&nbsp;7. [💻 Assembler CLI](#-assembler-cli)

##### *More*
### &nbsp;&nbsp;&nbsp;&nbsp;9. [📄 Contributing Guidelines](#-contributing-guidelines)

<br/>

---

<br/>

# 📌 Definition and Explanation
Assembly is basically the most basic programming language of a certain hardware. There's a very strong correspondence between the instructions in the language and the architecture's machine code instructions: every instruction in the language is a machine code instruction and vice-versa. It was created so that humans don't had to memorize the machine code instructions which are many numbers.

From the Wikipedia:
> In computer programming, assembly language, often abbreviated asm, is any low-level programming language in which there is a very strong correspondence between the instructions in the language and the architecture's machine code instructions.

Because of this strong correspondence, the translating process is called assembling instead of compiling, which is the same process but for high-end languages. Those languages do not have this strong correspondence that assembly languages have.

The core of the assembling process is to identify the assembly instructions and translate them to the circuit's instruction binary equivalent. Similarly, it also has to convert each variable to a memory address.

## Learn more

If you are interested in knowing more **how this process works** don't be afraid to read at the code.

If you are interested in knowing more **about the actual circuit** that runs the code you write, click [here](https://github.com/Open-Machine/Circuits/).

If you are interested in knowing more **about the Open-Computer project**, click [here](https://github.com/Open-Machine/README/).

<br/>

---

<br/>

# 🔢 Instructions
Let's take a close look at the instructions available. **Don't worry about syntax right now, we will talk about it later.**

**Warning** ⚠️: If you have never programmed in an assembly language (or with this assembly), please read this section and [🔀 Code Flow and Tips](#-code-flow-and-tips) in parallel. The *Code Flow and Tips* section will help you understand what to make with the instructions.

### Symbols Legend
Some symbols are used in the Instructions Table. Here you can see their meaning.
Symbol | Explanation
--- | ---
ACC | The ACC register
variable | A variable from the memory
label | Jump label
[ ] | "Value of"
```${varname}``` | Name of a defined variable
```${jumpLabel}``` | Jump label

### Instructions Table
Assembly Command | Short Instruction Description | Long Instruction Description | Short Param Description | Long Param Description
--- | --- | --- | --- | ---
```nop``` | - | This instruction doesn't perform any action | - | No parameter is required
```copy ${varname}``` | [ACC] = [variable] | A value from the memory is copied to the ACC register | variable | It's the name of the variable that will be used in the instruction
```store ${varname}``` | [variable] = [ACC] | The value from the ACC register is stored into memory | variable | It's the name of the variable that will be used in the instruction
```add ${varname}``` | [ACC] = [ACC] + [variable] | The sum of the value of the ACC register and a value from the memory is stored in the ACC register | variable | It's the name of the variable that will be used in the instruction
```sub ${varname}``` | [ACC] = [ACC] - [variable] | The difference between the value of the ACC register and a value from the memory is stored in the ACC register | variable | It's the name of the variable that will be used in the instruction
```input ${varname}``` | [variable] = input value | The input value is copied to the memory | variable | It's the name of the variable that will be used in the instruction
```output ${varname}``` | Output [variable] | Outputs a value from the memory into the circuit LEDs | variable | It's the name of the variable that will be used in the instruction
```kill``` | Finishes program | When this instruction is encountered, the program is finished and no more instructions will be executed | - | No parameter is required
```jmp ${jumpLabel}``` | Jump to EE | Jump to another line of code | label | The jump label the program will jump to
```jg ${jumpLabel}``` | Jump to EE if [ACC] > 0 | Jump to another line of code if the value of the ACC register is positive | label | The jump label the program will jump to if the condition is right
```je ${jumpLabel}``` | Jump to EE if [ACC] = 0 | Jump to another line of code if the value of the ACC register is zero | label | The jump label the program will jump to if the condition is right
```jl ${jumpLabel}``` | Jump to EE if [ACC] < 0 | Jump to another line of code if the value of the ACC register is negative | label | The jump label the program will jump to if the condition is right

<br/>

---

<br/>

# 🔀 Code Flow and Tips
This section will help you think more in an assembly way.

Because Open-Machine's Circuit only has very simple commands and very few registers, the way to think about your assembly code will be very different.

### [Click here](https://github.com/Open-Machine/Circuits/#-code-flow-and-tips) to go the section!

<br/>

---

<br/>

# 🔡 Code Syntax
**Warning** ⚠️: Assembly languages are specific to their hardware so remember that Open-Computer's Assembly may be different from other assembly languages.

Read the specifications below to learn the code syntax.

## General Definitions

### Tabs, spaces and case sensitivity
- Case sensitive;
- **Tabs and spaces** can be used interchangeably;
- Blank or empty lines won't be considered;
- **Numbers** can be written in hexadecimal in the form of ```0xff``` or in decimal as ```255```;
- Comments: characters written after a '#' character will be ignored in the assembly

### Naming Practices
- A **label name** should start with a lowercase letter and the rest of the name can have more letters and numbers;
- Every name should obey the following regex: ```[a-z][a-zA-Z0-9]*```;
- Snake-case is not allowed, use camel-case instead.

## The Two Parts
The file should be separated into 2 parts: the variable declaration and the actual code, as such:
```
@VAR
# you should declare your variables here

@CODE
# you should add the instructions and jump labels here
```

## Declaring Variables
The syntax of declaring variables is really easy: ```{variableName} = {int value}```.

For example: ```variable = 12``` or ```variable = 0x2```.

## Code

### Jump Label
- Definition: it marks the line for possible jumps to that line;
- Form: ```{labelName}:```
- Remember to follow the [naming practices](#naming-practices)

### Instruction line

***Definition***

An instruction line is a line that contains an instruction call.

***Components***
- ```instruction``` is the actual instruction that will be executed, it must be one of the following in the instruction table;
- ```arg``` can be a jump label or a number (depending on the instruction)

***Form***
- A instruction line should be in the following form ```{instruction} [arg]```;
- An instruction line should obey the following regex: ```^[\t ]*(((nop)|(copy)|(store)|(add)|(sub)|(input)|(output)|(kill)|(jmp)|(jg)|(je)|(jl))(([\t ]+[a-z][a-zA-Z0-9]*)|()))[\t ]*$```

***Instructions List***

Check out [here](#-instructions) the instruction table to know what instructions you can use and their parameters.

<br/>


---

<br/>

# ⌨️ Example
The following assembly code calculates the ```fibonacciNumberIndex```th fibonacci number.
```sh
@VAR
    fibonacciNumberIndex = 8

    oneVarConst = 1
    prev = 1
    current = 1
    auxNext = -1

@CODE
    copy fibonacciNumberIndex
    sub oneVarConst
    sub oneVarConst
    store fibonacciNumberIndex

    for:
        copy prev
        add current
        store auxNext

        copy current
        store prev

        copy auxNext
        store current

        copy fibonacciNumberIndex
        sub oneVarConst
        je end

        store fibonacciNumberIndex
        jmp for
    end:
    output current
    kill
```

<br/>

---

<br/>

# :arrow_forward: Setup and Run
These are the steps to setup and run Open-Computer's Assembler.

You can find more information about the assembler CLI [here](#-assembler-cli) and about running the circuit [here](https://github.com/Open-Machine/Circuits/#️-run).

## Setup
1. Build the GoLang project
	```sh
	./setup.sh
	```
2. Clone [Open-Computer's Circuit Repository](https://github.com/Open-Machine/Circuits/)

	You will need this repository to run the assembled program.

	If you have git installed in your terminal, run:
	```sh
	git clone https://github.com/Open-Machine/Circuits/
	```

## Run
There are two ways of running your application from the machine code generated by the ```assemble``` command.

### GUI Mode
In this mode, you will be able to see everything that is happening to the circuits in real time and interact with it by changing the inputs.

You can watch [this video](https://www.youtube.com/watch?v=NAITQqdOw7c) as an introduction to Logisim-Evolution, which is the program that we will be using to simulate the circuit.

1. Navigate to the Circuits repository
2. Start the circuit: follow the steps <a href="https://github.com/Open-Machine/Circuits/#i-start-the-circuit">to Start the Circuit</a>
3. Right click in the RAM and click "Load Image"
4. Select the assembled file

   You should select the file generated by the assemble program, not the file with the assembly code.

5. To run the program (start the clock simulation), follow the steps <a href="https://github.com/Open-Machine/Circuits/#iii-run-the-circuit">to Run the Circuit</a>

## 💻 Assembler CLI
You can use the flag ```--help``` to see all the options.

### Run
```sh
./assembler run --help
```
```
usage: assembly run <file-name> [<number-format>]

Run machine code

Flags:
    --help  Show context-sensitive help (also try --help-long and
            --help-man).

Args:
  <file-name>        Provide the name of file with the assembly or machine code
  [<number-format>]  Provide the format for the numbers printed
```

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

## Test
```sh
go test ./...
cd go_scripts/format_circuit_output
go test ./...
```

<br/>

---

<br/>

# 📄 Contributing Guidelines
Check out the contributing guidelines [here](https://github.com/Open-Machine/Assembler/blob/stable/CONTRIBUTING.md).
