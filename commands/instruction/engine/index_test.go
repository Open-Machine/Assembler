package engine

// TODO

// func TestGetNoParam(t *testing.T) {
// 	got, err := getParamNoParam("nop", []string{"nop"})
// 	if !(err == nil && !got.IsStr && got.Num == 0) {
// 		t.Errorf("Wrong")
// 	}
// }

// func TestGetSecondParamAsInt(t *testing.T) {
// 	var tests = []struct {
// 		line       string
// 		expected   *data.InstructionParameter
// 		expectsErr bool
// 	}{
// 		// Decimal Number
// 		{"mov 1", newCmdIntParam(1), false},
// 		// Hexadecimal Number
// 		{"mov 0x1a", newCmdIntParam(26), false},
// 		{"mov 0x001", newCmdIntParam(1), false},
// 		{"mov 0x0f", newCmdIntParam(15), false},
// 		{"mov 0xff", newCmdIntParam(255), false},
// 		{"mov 0x0ff", newCmdIntParam(255), false},
// 		// Variable
// 		{"mov 0xx0ff", nil, true},
// 		{"mov x1", nil, true},
// 		{"mov 0x1g", nil, true},
// 		{"mov 1a", nil, true},
// 		// Words
// 		{"mov", nil, true},
// 		{"mov a b", nil, true},
// 		{"mov a b c", nil, true},
// 	}

// 	for i, test := range tests {
// 		arrayWords := strings.Split(test.line, " ")
// 		got, err := getSecondWord(arrayWords[0], arrayWords, false)
// 		gotError := err != nil

// 		if test.expectsErr != gotError {
// 			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotError)
// 		}

// 		if !helper.SafeIsEqualInstructionParamPointer(test.expected, got) {
// 			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expected, got)
// 		}

// 		if got != nil && got.IsStr {
// 			t.Errorf("[%d] Expecting int parameter", i)
// 		}
// 	}
// }

// func TestGetSecondParamAsIntOrString(t *testing.T) {
// 	var tests = []struct {
// 		line       string
// 		expected   *data.InstructionParameter
// 		expectsErr bool
// 	}{
// 		// Decimal Number
// 		{"jmp 1", newCmdIntParam(1), false},
// 		// Hexadecimal Number
// 		{"jmp 0x001", newCmdIntParam(1), false},
// 		{"jmp 0x0f", newCmdIntParam(15), false},
// 		{"jmp 0xff", newCmdIntParam(255), false},
// 		{"jmp 0x0ff", newCmdIntParam(255), false},
// 		// Variable
// 		{"jmp a8", newCmdStringParam("a8"), false},
// 		{"jmp x1", newCmdStringParam("x1"), false},
// 		{"jmp a1", newCmdStringParam("a1"), false},
// 		// Errors 1 param
// 		{"jmp 0xx0ff", nil, true},
// 		{"jmp 0x1g", nil, true},
// 		{"jmp 1a", nil, true},
// 		// Erros amnt params
// 		{"jmp", nil, true},
// 		{"jmp a b", nil, true},
// 		{"jmp a b c", nil, true},
// 	}

// 	for i, test := range tests {
// 		arrayWords := strings.Split(test.line, " ")
// 		got, err := getSecondWord(arrayWords[0], arrayWords, true)
// 		gotError := err != nil

// 		if test.expectsErr != gotError {
// 			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotError)
// 		}

// 		if !helper.SafeIsEqualInstructionParamPointer(test.expected, got) {
// 			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expected, got)
// 		}

// 		if got != nil && test.expected != nil && got.IsStr != test.expected.IsStr {
// 			t.Errorf("[%d] Expected IsStr: %t, Got IsStr: %t", i, test.expected.IsStr, got.IsStr)
// 		}
// 	}
// }
// func newCmdIntParam(num int) *data.InstructionParameter {
// 	param := data.NewIntParam(num)
// 	return &param
// }
// func newCmdStringParam(str string) *data.InstructionParameter {
// 	param := data.NewStringParam(str)
// 	return &param
// }

// func TestAssembleInstruction(t *testing.T) {
// 	if len(instructions) != 12 {
// 		t.Errorf("Tests were not updated")
// 	}

// 	var tests = []struct {
// 		line       string
// 		expected   *data.Instruction
// 		expectsErr bool
// 	}{
// 		// Success Number
// 		{"nop", getInstruction(0x0, 0), false},
// 		{"copy 0x10", getInstruction(0x1, 16), false},
// 		{"store 0x10", getInstruction(0x2, 16), false},
// 		{"add 10", getInstruction(0x3, 10), false},
// 		{"sub 10", getInstruction(0x4, 10), false},
// 		{"input 7", getInstruction(0x7, 7), false},
// 		{"output 8", getInstruction(0x8, 8), false},
// 		{"kill", getInstruction(0x9, 0), false},
// 		{"jmp 0x8", getInstruction(0xA, 8), false},
// 		{"jg 0x8", getInstruction(0xB, 8), false},
// 		{"je 0x8", getInstruction(0xD, 8), false},
// 		{"jl 0x8", getInstruction(0xF, 8), false},
// 		// Success Label
// 		{"jmp label", getInstructionStr(0xA, "label"), false},
// 		{"jg label", getInstructionStr(0xB, "label"), false},
// 		{"je label", getInstructionStr(0xD, "label"), false},
// 		{"jl label", getInstructionStr(0xF, "label"), false},
// 		// Fail: Wrong Instruction
// 		{"nope", nil, true},
// 		// Fail: No label as param
// 		{"copy label", nil, true},
// 		{"store label", nil, true},
// 		{"add label", nil, true},
// 		{"sub label", nil, true},
// 		{"input label", nil, true},
// 		{"output label", nil, true},
// 		// Fail: Wrong param
// 		{"add 1x10", nil, true},
// 		// Fail: Amnt params
// 		{"kill 0", nil, true},
// 		{"output", nil, true},
// 		{"output 8 1", nil, true},
// 	}

// 	for i, test := range tests {
// 		got, err := AssembleInstruction(test.line)
// 		gotError := err != nil

// 		if test.expectsErr != gotError {
// 			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotError)
// 		}
// 		if !helper.SafeIsEqualInstructionPointer(test.expected, got) {
// 			t.Errorf("Instruction expected is: %v, Got expected is: %v", test.expected, got)
// 		}
// 	}
// }
// func getInstruction(code int, param int) *data.Instruction {
// 	cmd, _ := data.NewInstruction(code, data.NewIntParam(param))
// 	return cmd
// }
// func getInstructionStr(code int, param string) *data.Instruction {
// 	cmd, _ := data.NewInstruction(code, data.NewStringParam(param))
// 	return cmd
// }

// func TestGetInstructionParams(t *testing.T) {
// 	var tests = []struct {
// 		param    []string
// 		expected []string
// 	}{
// 		{
// 			[]string{"mov", "0x1", "1", "label"},
// 			[]string{"0x1", "1", "label"},
// 		},
// 		{
// 			[]string{"cp", "to", "here"},
// 			[]string{"to", "here"},
// 		},
// 		{
// 			[]string{"cp"},
// 			[]string{},
// 		},
// 	}

// 	for i, test := range tests {
// 		got := getInstructionParams(test.param)
// 		if !reflect.DeepEqual(got, test.expected) {
// 			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expected, got)
// 		}
// 	}
// }

// func TestCheckAmntWords(t *testing.T) {
// 	var tests = []struct {
// 		amntParamsExpected int
// 		instructionName    string
// 		words              []string
// 		expectedErr        bool
// 	}{
// 		// Success: no param
// 		{
// 			amntParamsExpected: 0,
// 			instructionName:    "mov",
// 			words:              []string{},
// 			expectedErr:        false,
// 		},
// 		// Success: one param
// 		{
// 			amntParamsExpected: 1,
// 			instructionName:    "mov",
// 			words:              []string{"hello"},
// 			expectedErr:        false,
// 		},
// 		// Success: two param
// 		{
// 			amntParamsExpected: 2,
// 			instructionName:    "mov",
// 			words:              []string{"hello", "there"},
// 			expectedErr:        false,
// 		},
// 		// Fail: expected more (no param)
// 		{
// 			amntParamsExpected: 1,
// 			instructionName:    "mov",
// 			words:              []string{},
// 			expectedErr:        true,
// 		},
// 		// Fail: expected more (more params)
// 		{
// 			amntParamsExpected: 5,
// 			instructionName:    "mov",
// 			words:              []string{"hello", "there", "now"},
// 			expectedErr:        true,
// 		},
// 		// Fail: expected less
// 		{
// 			amntParamsExpected: 0,
// 			instructionName:    "mov",
// 			words:              []string{"hello"},
// 			expectedErr:        true,
// 		},
// 	}

// 	for i, test := range tests {
// 		err := checkAmntWords(test.amntParamsExpected, test.instructionName, test.words)
// 		gotErr := err != nil

// 		if gotErr != test.expectedErr {
// 			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectedError, gotErr)
// 		}
// 	}
// }

// func TestAddInstruction(t *testing.T) {
// 	var tests = []struct {
// 		code            int
// 		param           data.InstructionParameter
// 		program         data.Program
// 		expectedError   bool
// 		expectedProgram data.Program
// 	}{
// 		// Success
// 		{
// 			code:  5,
// 			param: 1000,
// 			program: newProgramPointer([]data.Instruction{
// 				newInstruction(1, 2),
// 				newInstruction(3, 4),
// 			}),
// 			expectedError: false,
// 			expectedProgram: newProgramPointer([]data.Instruction{
// 				newInstruction(1, 2),
// 				newInstruction(3, 4),
// 				newInstruction(5, 1000),
// 			}),
// 		},
// 		// Fail
// 		{
// 			code:  1000,
// 			param: 200,
// 			program: newProgramPointer([]data.Instruction{
// 				newInstruction(1, 2),
// 				newInstruction(3, 4),
// 			}),
// 			expectedError: true,
// 			expectedProgram: newProgramPointer([]data.Instruction{
// 				newInstruction(1, 2),
// 				newInstruction(3, 4),
// 			}),
// 		},
// 	}

// 	for i, test := range tests {
// 		err := addInstruction(test.code, test.param, test.program)
// 		gotErr := err != nil

// 		if gotErr != test.expectedError {
// 			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectedError, gotErr)
// 		}
// 		if test.program != test.expectedProgram {
// 			t.Errorf("[%d] Expected program: %v, Got program: %v", i, test.expectedProgram, test.program)
// 		}
// 	}
// }

// func newInstruction(code int, numParam int) data.Instruction {
// 	config.Testing = true
// 	param := data.NewParamTest(numParam, "", true)
// 	instruc, _ := data.NewInstruction(code, param)
// 	return instruc
// }

// func newProgramPointer(instructions []data.Instruction) *data.Program {
// 	prog := data.ProgramFromInstructionsAndLabels(instructions, map[string]int{})
// 	return &prog
// }
